package prestashop

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	maintainAccessTokenBuffer     = 300
	authenticationRetryAttempts   = 3
	authenticationRetryHoldMillis = 2000
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func (client *Client) AuthenticateWithToken(accessToken string) {
	client.auth.HeaderName = defaultAuthHeaderName
	client.auth.AccessToken = "Bearer " + accessToken
}

func (client *Client) Authenticate(clientName string, clientID string, clientSecret string, maintainToken bool) (*AccessTokenResponse, error) {
	accessToken, _, err := client.GenerateAccessToken(clientName, clientID, clientSecret)

	if err != nil {
		return accessToken, err
	}

	if accessToken == nil {
		return accessToken, errors.New("access token missing")
	}

	client.AuthenticateWithToken(accessToken.AccessToken)

	if maintainToken {
		go client.maintainAccessToken(clientName, clientID, clientSecret, accessToken.ExpiresIn)
	}

	return accessToken, nil
}

func (client *Client) maintainAccessToken(clientName string, clientID string, clientSecret string, expiresIn int) {
	// Initialize maintain token if not already initialized
	if client.maintainToken == nil {
		client.maintainToken = &maintainAccessToken{
			Attempt:   0,
			LastError: nil,
			Stopped:   false,
		}
	}

	// Wait before renewing access token using expire and maintain buffer
	time.Sleep(time.Second * time.Duration(expiresIn-maintainAccessTokenBuffer))

	for client.maintainToken.Attempt < authenticationRetryAttempts {
		// Hold before this attempt? (ie. not first attempt)
		if client.maintainToken.Attempt > 0 {
			time.Sleep(authenticationRetryHoldMillis * time.Millisecond)
		}

		// Increment attempt counter
		client.maintainToken.Attempt++

		// Re-authenticate client
		_, err := client.Authenticate(clientName, clientID, clientSecret, true)

		// Re-authentication successfull, stop there.
		if err == nil {
			// Reset attempt counter and last error
			client.maintainToken.Attempt = 0
			client.maintainToken.LastError = nil

			break
		}

		// Reached max attempts, set stopped flag.
		if client.maintainToken.Attempt >= authenticationRetryAttempts {
			client.maintainToken.Stopped = true
		}

		client.maintainToken.LastError = err
	}
}

func (client *Client) GenerateAccessToken(clientName string, clientID string, clientSecret string) (*AccessTokenResponse, *http.Response, error) {
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("grant_type", "client_credentials")
	data.Set("audience", strings.Join([]string{defaultRestEndpointURL, " ", "tech-vendor/" + clientName}, ""))

	req, err := http.NewRequest("POST", defaultRestOAuthEndpointURL, bytes.NewBufferString(data.Encode()))

	if err != nil {
		return nil, nil, err
	}

	req.Header.Add(defaultAuthHeaderName, generateBasicBase64Token(clientID, clientSecret))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", acceptedContentType)
	req.Header.Add("User-Agent", userAgent)

	accessToken := new(AccessTokenResponse)
	resp, err := client.Do(req, accessToken)

	if err != nil {
		return nil, resp, err
	}

	return accessToken, resp, err
}

func generateBasicBase64Token(clientID string, clientSecret string) string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(strings.Join([]string{clientID, clientSecret}, ":"))))
}

// Returns stopped flag (will not attempt to re-authenticate), attempt counter and last error.
func (client *Client) FetchMaintainTokenStopped() (bool, int, error) {
	if client.maintainToken == nil {
		return true, -1, nil
	}

	return client.maintainToken.Stopped, client.maintainToken.Attempt, client.maintainToken.LastError
}
