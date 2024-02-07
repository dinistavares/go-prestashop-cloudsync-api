package prestashop

import (
	"fmt"
	"net/http"
)

// CarrierTax service
type CarrierTaxesService service

type CarrierTaxesResponse struct {
	CarrierTax   *CarrierTax   `json:"item,omitempty"`
	CarrierTaxes *[]CarrierTax `json:"items,omitempty"`
	GenericListResponse
}

type CarrierTax struct {
	ID                   string `json:"id,omitempty"`
	ShopID               string `json:"shopId,omitempty"`
	CreatedAt            string `json:"createdAt,omitempty"`
	UpdatedAt            string `json:"updatedAt,omitempty"`
	LastSyncedAt         string `json:"lastSyncedAt,omitempty"`
	IDReference          string `json:"idReference,omitempty"`
	IDCarrierTax         string `json:"idCarrierTax,omitempty"`
	CountryID            string `json:"countryId,omitempty"`
	StateIds             string `json:"stateIds,omitempty"`
	TaxRate              int    `json:"taxRate,omitempty"`
	CarrierTaxesUniqueID string `json:"carrierTaxesUniqueId,omitempty"`
	IDRange              int    `json:"idRange,omitempty"`
	IDZone               int    `json:"idZone,omitempty"`
}

type CarrierTaxesListParams struct {
	IdCarrierTax         string  `url:"idCarrierTax,omitempty"`
	IdRange              string  `url:"idRange,omitempty"`
	IdReference          string  `url:"idReference,omitempty"`
	IdZone               string  `url:"idZone,omitempty"`
	CountryId            string  `url:"countryId,omitempty"`
	StateIds             string  `url:"stateIds,omitempty"`
	CarrierTaxesUniqueId string  `url:"carrierTaxesUniqueId,omitempty"`
	TaxRate              float64 `url:"taxRate,omitempty"`
	TaxRateStart         float64 `url:"taxRateStart,omitempty"`
	TaxRateEnd           float64 `url:"taxRateEnd,omitempty"`
	GenericListParams
}

// List carrier taxes. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/CarrierTaxes_getPaginatedItems
func (service *CarrierTaxesService) List(shopID string, opts *CarrierTaxesListParams) (*CarrierTaxesResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/carrier-taxes", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	carrierTaxes := new(CarrierTaxesResponse)
	response, err := service.client.Do(req, carrierTaxes)

	if err != nil {
		return nil, response, err
	}

	return carrierTaxes, response, nil
}

// Get a carrier tax. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/CarrierTaxes_getSingleItem
func (service *CarrierTaxesService) Get(shopID string, shopContentId string) (*CarrierTaxesResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/carrier-taxes/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	carrierDetail := new(CarrierTaxesResponse)
	response, err := service.client.Do(req, carrierDetail)

	if err != nil {
		return nil, response, err
	}

	return carrierDetail, response, nil
}
