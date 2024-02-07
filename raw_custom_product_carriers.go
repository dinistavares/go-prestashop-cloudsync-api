package prestashop

import (
	"fmt"
	"net/http"
)

// Custom Product Carrier service
type CustomProductCarriersService service

type CustomProductCarriersResponse struct {
	CustomProductCarrier  *CustomProductCarrier   `json:"item,omitempty"`
	CustomProductCarriers *[]CustomProductCarrier `json:"items,omitempty"`
	GenericListResponse
}

type CustomProductCarrier struct {
	ID                     string `json:"id,omitempty"`
	ShopID                 string `json:"shopId,omitempty"`
	CreatedAt              string `json:"createdAt,omitempty"`
	UpdatedAt              string `json:"updatedAt,omitempty"`
	LastSyncedAt           string `json:"lastSyncedAt,omitempty"`
	IDCustomProductCarrier string `json:"idCustomProductCarrier,omitempty"`
	IDCarrierReference     int    `json:"idCarrierReference,omitempty"`
	IDProduct              int    `json:"idProduct,omitempty"`
}

type CustomProductCarrierListParams struct {
	IdCarrierReference       int     `url:"idCarrierReference,omitempty"`
	IdCarrierReferenceStart  int     `url:"idCarrierReferenceStart,omitempty"`
	IdCarrierReferenceEnd    int     `url:"idCarrierReferenceEnd,omitempty"`
	IdProduct                int     `url:"idProduct,omitempty"`
	IdProductStart           int     `url:"idProductStart,omitempty"`
	IdProductEnd             int     `url:"idProductEnd,omitempty"`
	IdCustomProductCarrier   string  `url:"idCustomProductCarrier,omitempty"`
	GenericListParams
}


// List custom Product Carriers. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/CustomProductCarriers_getPaginatedItems
func (service *CustomProductCarriersService) List(shopID string, opts *CustomProductCarrierListParams) (*CustomProductCarriersResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/custom-product-carriers", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	customProductCarriers := new(CustomProductCarriersResponse)
	response, err := service.client.Do(req, customProductCarriers)

	if err != nil {
		return nil, response, err
	}

	return customProductCarriers, response, nil
}

// Get a custom Product Carrier. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/CustomProductCarriers_getSingleItem
func (service *CustomProductCarriersService) Get(shopID string, shopContentId string) (*CustomProductCarriersResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/custom-product-carriers/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	customer := new(CustomProductCarriersResponse)
	response, err := service.client.Do(req, customer)

	if err != nil {
		return nil, response, err
	}

	return customer, response, nil
}
