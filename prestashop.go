package prestashop

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	libraryVersion               = "1.5"
	defaultRestEndpointVersion   = "v1"
	defaultAuthHeaderName        = "Authorization"
	defaultRestEndpointURL       = "https://api.cloudsync.prestashop.com"
	defaultRestOAuthEndpointURL  = "https://oauth.prestashop.com/oauth2/token"
	acceptedContentType          = "application/json"
	userAgent                    = "go-prestashop-cloudsync-api/" + libraryVersion
	clientRequestRetryAttempts   = 2
	clientRequestRetryHoldMillis = 1000
)

var (
	errorDoAllAttemptsExhausted = errors.New("all request attempts were exhausted")
	errorDoAttemptNilRequest    = errors.New("request could not be constructed")

	// Resourece type 'raw'
	ResourceTypeRaw ResourceType = "raw"

	// Resourece type 'reporting'
	ResourceTypeReporting ResourceType = "reporting"

	// Resourece type 'sync'
	ResourceTypeSync ResourceType = "sync"
)

type ResourceType string

type ClientConfig struct {
	HttpClient          *http.Client
	RestEndpointURL     string
	RestEndpointVersion string
}

type auth struct {
	HeaderName  string
	AccessToken string
}

type Client struct {
	config        *ClientConfig
	client        *http.Client
	auth          *auth
	baseURL       *url.URL
	maintainToken *maintainAccessToken

	RawApi       *RawApi
	ReportingApi *ReportingApi
	SyncApi      *SyncApi
}

type maintainAccessToken struct {
	Attempt   int
	LastError error
	Stopped   bool
}

type RawApi struct {
	Bundles               *BundlesService
	Carriers              *CarriersService
	CarrierDetails        *CarrierDetailsService
	CarrierTaxes          *CarrierTaxesService
	Carts                 *CartsService
	CartProducts          *CartProductsService
	CartRules             *CartRulesService
	Categories            *CategoriesService
	Currencies            *CurrenciesService
	CustomProductCarriers *CustomProductCarriersService
	Customers             *CustomersService
	Orders                *OrdersService
	OrderDetails          *OrderDetailsService
	OrderStatusHistory    *OrderStatusHistoryService
	Products              *ProductsService
	Shop                  *ShopService
}

type SyncApi struct {
	Sync *TriggerSyncService
}

type ReportingApi struct {
	SyncDetails *SyncDetailsService
	SyncErrors  *SyncErrorsService
}

type service struct {
	client *Client
}

type errorResponse struct {
	Response *http.Response

	Code             string      `json:"code"`
	Message          string      `json:"message"`
	ErrorDescription string      `json:"error_description"`
	StatusText       interface{} `json:"statusText"`
	Data             ErrorData   `json:"data"`
}

type ErrorData struct {
	Status int `json:"status"`
}

func (response *errorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v %v",
		response.Response.Request.Method, response.Response.Request.URL,
		response.Response.StatusCode, response.Message, response.ErrorDescription,
		response.StatusText)
}

func New() (*Client, error) {
	config := ClientConfig{
		HttpClient: http.DefaultClient,
	}

	config.HttpClient = http.DefaultClient
	config.RestEndpointURL = defaultRestEndpointURL
	config.RestEndpointVersion = defaultRestEndpointVersion

	// Create client
	baseURL, err := url.Parse(config.RestEndpointURL)

	if err != nil {
		return nil, err
	}

	client := &Client{config: &config, client: config.HttpClient, auth: &auth{}, baseURL: baseURL}

	// Map services
	client.RawApi = &RawApi{
		Bundles:               &BundlesService{client: client},
		Carriers:              &CarriersService{client: client},
		CarrierDetails:        &CarrierDetailsService{client: client},
		CarrierTaxes:          &CarrierTaxesService{client: client},
		Carts:                 &CartsService{client: client},
		CartProducts:          &CartProductsService{client: client},
		CartRules:             &CartRulesService{client: client},
		Categories:            &CategoriesService{client: client},
		Currencies:            &CurrenciesService{client: client},
		CustomProductCarriers: &CustomProductCarriersService{client: client},
		Customers:             &CustomersService{client: client},
		Orders:                &OrdersService{client: client},
		OrderDetails:          &OrderDetailsService{client: client},
		OrderStatusHistory:    &OrderStatusHistoryService{client: client},
		Products:              &ProductsService{client: client},
		Shop:                  &ShopService{client: client},
	}

	client.SyncApi = &SyncApi{
		Sync: &TriggerSyncService{client: client},
	}

	client.ReportingApi = &ReportingApi{
		SyncDetails: &SyncDetailsService{client: client},
		SyncErrors:  &SyncErrorsService{client: client},
	}

	return client, nil
}

// NewRequest creates an API request
func (client *Client) NewRequest(method, urlStr string, opts interface{}, body interface{}) (*http.Request, error) {
	// Append Query Params to URL
	if opts != nil {
		queryParams, err := query.Values(opts)

		if err != nil {
			return nil, err
		}

		rawQuery := queryParams.Encode()

		if rawQuery != "" {
			urlStr += "?" + rawQuery
		}
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	url := client.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)

		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.auth.HeaderName, client.auth.AccessToken)
	req.Header.Add("Accept", acceptedContentType)
	req.Header.Add("Content-type", acceptedContentType)
	req.Header.Add("User-Agent", userAgent)

	return req, nil
}

// Do sends an API request
func (client *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	var lastErr error

	attempts := 0

	for attempts < clientRequestRetryAttempts {
		// Hold before this attempt? (ie. not first attempt)
		if attempts > 0 {
			time.Sleep(clientRequestRetryHoldMillis * time.Millisecond)
		}

		// Dispatch request attempt
		attempts++
		resp, shouldRetry, err := client.doAttempt(req, v)

		// Return response straight away? (we are done)
		if !shouldRetry {
			return resp, err
		}

		// Should retry: store last error (we are not done)
		lastErr = err
	}

	// Set default error? (all attempts failed, but no error is set)
	if lastErr == nil {
		lastErr = errorDoAllAttemptsExhausted
	}

	// All attempts failed, return last attempt error
	return nil, lastErr
}

func (client *Client) doAttempt(req *http.Request, v interface{}) (*http.Response, bool, error) {
	if req == nil {
		return nil, false, errorDoAttemptNilRequest
	}

	resp, err := client.client.Do(req)

	if checkRequestRetry(resp, err) {
		return nil, true, err
	}

	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return resp, false, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}

	return resp, false, err
}

// checkRequestRetry checks if should retry request
func checkRequestRetry(response *http.Response, err error) bool {
	// Low-level error, or response status is a server error? (HTTP 5xx)
	if err != nil || response.StatusCode >= 500 {
		return true
	}

	// No low-level error (should not retry)
	return false
}

// checkResponse checks response for errors
func checkResponse(response *http.Response) error {
	// No error in response? (HTTP 2xx)
	if code := response.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	// Map response error data (eg. HTTP 4xx)
	errorResponse := &errorResponse{Response: response}

	data, err := io.ReadAll(response.Body)
	if err == nil && data != nil {
		_ = json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}
