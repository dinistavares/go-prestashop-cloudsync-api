package prestashop

import (
	"fmt"
	"net/http"
	"time"
)

// SyncErrors service
type SyncErrorsService service

type SyncErrorsResponse struct {
	Start  *time.Time   `json:"start,omitempty"`
	End    *time.Time   `json:"end,omitempty"`
	Errors *[]SyncError `json:"errors,omitempty"`
}

type SyncError struct {
	Code         int    `json:"code,omitempty"`
	Distribution int    `json:"distribution,omitempty"`
	Description  string `json:"description,omitempty"`
}

type SyncErrorsParams struct {
	ShopID string `url:"shopId,omitempty"`
	Start  string `url:"start,omitempty"`
	End    string `url:"end,omitempty"`
}

// Get all errors by partner. Reference: https://docs.cloudsync.prestashop.com/api-doc/reporting-api#/operations/VendorErrorCodeController_getVendorErrorsCode
func (service *SyncErrorsService) GetSyncErrorsByPartner(vendorID string, opts *SyncErrorsParams) (*SyncErrorsResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/vendor/%s/errors", service.client.getResourceTypeReporting(), vendorID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	responseBody := new(SyncErrorsResponse)
	response, err := service.client.Do(req, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

// Get all errors by shop. Reference: https://docs.cloudsync.prestashop.com/api-doc/reporting-api#/operations/ShopErrorCodeController_getShopErrorsCode
func (service *SyncErrorsService) GetSyncErrorsByShop(shopID string, opts *SyncErrorsParams) (*SyncErrorsResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/shop/%s/errors", service.client.getResourceTypeReporting(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	responseBody := new(SyncErrorsResponse)
	response, err := service.client.Do(req, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}
