package prestashop

import (
	"fmt"
	"net/http"
)

// CarrierDetail service
type CarrierDetailsService service

type CarrierDetailsResponse struct {
	CarrierDetail  CarrierDetail   `json:"item,omitempty"`
	CarrierDetails []CarrierDetail `json:"items,omitempty"`
	GenericListResponse
}

type CarrierDetail struct {
	ID                     string  `json:"id,omitempty"`
	ShopID                 string  `json:"shopId,omitempty"`
	CreatedAt              string  `json:"createdAt,omitempty"`
	UpdatedAt              string  `json:"updatedAt,omitempty"`
	LastSyncedAt           string  `json:"lastSyncedAt,omitempty"`
	IDReference            string  `json:"idReference,omitempty"`
	ShippingMethod         string  `json:"shippingMethod,omitempty"`
	Delimiter1             int     `json:"delimiter1,omitempty"`
	Delimiter2             int     `json:"delimiter2,omitempty"`
	CountryIds             string  `json:"countryIds,omitempty"`
	StateIds               string  `json:"stateIds,omitempty"`
	Price                  float64 `json:"price,omitempty"`
	IDCarrierDetail        string  `json:"idCarrierDetail,omitempty"`
	CarrierDetailsUniqueID string  `json:"carrierDetailsUniqueId,omitempty"`
	IDRange                string  `json:"idRange,omitempty"`
	IDZone                 string  `json:"idZone,omitempty"`
}

type CarrierDetailListParams struct {
	CarrierDetailsUniqueId string  `url:"carrierDetailsUniqueId,omitempty"`
	CountryIds             string  `url:"countryIds,omitempty"`
	IdCarrierDetail        string  `url:"idCarrierDetail,omitempty"`
	IdRange                string  `url:"idRange,omitempty"`
	IdReference            string  `url:"idReference,omitempty"`
	IdZone                 string  `url:"idZone,omitempty"`
	Price                  float64 `url:"price,omitempty"`
	PriceStart             float64 `url:"priceStart,omitempty"`
	PriceEnd               float64 `url:"priceEnd,omitempty"`
	ShippingMethod         string  `url:"shippingMethod,omitempty"`
	StateIds               string  `url:"stateIds,omitempty"`
	Delimiter1             int     `url:"delimiter1,omitempty"`
	Delimiter1Start        int     `url:"delimiter1Start,omitempty"`
	Delimiter1End          int     `url:"delimiter1End,omitempty"` // Note: Adjusted typo from `delimiter1SEnd` to `delimiter1End`
	Delimiter2             int     `url:"delimiter2,omitempty"`
	Delimiter2Start        int     `url:"delimiter2Start,omitempty"`
	Delimiter2End          int     `url:"delimiter2End,omitempty"`
	GenericListParams
}

// List carrier details. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/CarrierDetails_getPaginatedItems
func (service *CarrierDetailsService) List(shopID string, opts *CarrierDetailListParams) (*CarrierDetailsResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/carrier-details", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	carrierDetails := new(CarrierDetailsResponse)
	response, err := service.client.Do(req, carrierDetails)

	if err != nil {
		return nil, response, err
	}

	return carrierDetails, response, nil
}

// Get a carrier detail. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/CarrierDetails_getSingleItem
func (service *CarrierDetailsService) Get(shopID string, shopContentId string) (*CarrierDetailsResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/carrier-details/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	carrierDetail := new(CarrierDetailsResponse)
	response, err := service.client.Do(req, carrierDetail)

	if err != nil {
		return nil, response, err
	}

	return carrierDetail, response, nil
}
