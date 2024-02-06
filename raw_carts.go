package prestashop

import (
	"fmt"
	"net/http"
)

// Cart service
type CartsService service

type CartsResponse struct {
	Cart  *Cart   `json:"item,omitempty"`
	Carts *[]Cart `json:"items,omitempty"`
	GenericListResponse
}

type Cart struct {
	ID           string `json:"id,omitempty"`
	ShopID       string `json:"shopId,omitempty"`
	IDCart       string `json:"idCart,omitempty"`
	UpdatedAt    string `json:"updatedAt,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
	LastSyncedAt string `json:"lastSyncedAt,omitempty"`
}

type CartListParams struct {
	IdCart string `url:"idCart,omitempty"`
	GenericListParams
}

// List carts. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Carts_getPaginatedItems
func (service *CartsService) List(shopID string, opts *CartListParams) (*CartsResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/carts", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	carts := new(CartsResponse)
	response, err := service.client.Do(req, carts)

	if err != nil {
		return nil, response, err
	}

	return carts, response, nil
}

// Get a cart. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Carts_getSingleItem
func (service *CartsService) Get(shopID string, shopContentId string) (*CartsResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/carts/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	cart := new(CartsResponse)
	response, err := service.client.Do(req, cart)

	if err != nil {
		return nil, response, err
	}

	return cart, response, nil
}
