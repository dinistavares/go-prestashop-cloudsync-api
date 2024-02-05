package prestashop

import (
	"fmt"
	"net/http"
)

// Trigger service
type TriggerSyncService service

type TriggerResponse struct {
	Result string `json:"result,omitempty"`
}

type SyncTrigger struct {
	ShopID       string   `json:"shopId,omitempty"`
	ShopContents []string `json:"shopContents,omitempty"`
}

// Trigger sync. Reference: https://docs.cloudsync.prestashop.com/api-doc/sync-api#/operations/AuthSyncController_triggerSyncAuth
func (service *TriggerSyncService) TriggerSync(trigger *SyncTrigger) (*TriggerResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/sync/trigger", service.client.getResourceTypeSync())

	req, _ := service.client.NewRequest("POST", _url, nil, trigger)

	responseBody := new(TriggerResponse)
	response, err := service.client.Do(req, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

// Trigger sync. Reference: https://docs.cloudsync.prestashop.com/api-doc/sync-api#/operations/AuthSyncController_triggerFullSyncAuth
func (service *TriggerSyncService) TriggerFullSync(trigger *SyncTrigger) (*TriggerResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/sync/trigger-full", service.client.getResourceTypeSync())

	req, _ := service.client.NewRequest("POST", _url, nil, trigger)

	responseBody := new(TriggerResponse)
	response, err := service.client.Do(req, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}
