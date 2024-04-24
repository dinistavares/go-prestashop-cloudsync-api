package prestashop

import (
	"fmt"
	"net/http"
	"time"
)

// SyncErrors service
type SyncDetailsService service

type SyncDetailsSetupResponse struct {
	TargetURL     string    `json:"targetUrl,omitempty"`
	ShopContent   *[]string `json:"shopContent,omitempty"`
	ShopLanguages *[]string `json:"shopLanguages,omitempty"`
}

type SyncDetailsSummaryResponse struct {
	FirstSyncFinished *time.Time `json:"firstSyncFinishedAt,omitempty"`
	LastSyncFinished  *time.Time `json:"lastSyncFinishedAt,omitempty"`
	ShopContent       string     `json:"shopContent,omitempty"`
}

type SyncDetailsOnboardingResponse struct {
	FirstOnboarding *time.Time `json:"firstOnboarding,omitempty"`
	LastOnboarding  *time.Time `json:"lastOnboarding,omitempty"`
	TargetUrl       string     `json:"targetUrl,omitempty"`
}

type SyncDetailsParams struct {
	ShopContent string `url:"shopContent,omitempty"`
}

// Get all errors by partner. Reference: https://docs.cloudsync.prestashop.com/api-doc/reporting-api#/operations/ShopSyncSetupController_getShopSyncSetup
func (service *SyncDetailsService) GetShopSyncSetup(shopID string) (*SyncDetailsSetupResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/shop-sync-setup/%s", service.client.getResourceTypeReporting(), shopID)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	responseBody := new(SyncDetailsSetupResponse)
	response, err := service.client.Do(req, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

// Get all errors by partner. Reference: https://docs.cloudsync.prestashop.com/api-doc/reporting-api#/operations/ShopSyncSummaryController_getSyncSummaryByShopId
func (service *SyncDetailsService) GetShopSyncSummary(shopID string, opts *SyncDetailsParams) (*SyncDetailsSummaryResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/shop-sync-summary/%s", service.client.getResourceTypeReporting(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	responseBody := new(SyncDetailsSummaryResponse)
	response, err := service.client.Do(req, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

// Get all errors by partner. Reference: https://docs.cloudsync.prestashop.com/api-doc/reporting-api#/operations/ShopOnboardingController_getOnboardingInfoByShopId
func (service *SyncDetailsService) GetShopSyncOnboarding(shopID string) (*SyncDetailsOnboardingResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/shop-sync-onboarding/%s", service.client.getResourceTypeReporting(), shopID)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	responseBody := new(SyncDetailsOnboardingResponse)
	response, err := service.client.Do(req, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}
