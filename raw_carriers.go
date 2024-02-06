package prestashop

import (
	"fmt"
	"net/http"
)

// Carrier service
type CarriersService service

type CarriersResponse struct {
	Carrier  Carrier   `json:"item,omitempty"`
	Carriers []Carrier `json:"items,omitempty"`
	GenericListResponse
}

type Carrier struct {
	ID                           string `json:"id,omitempty"`
	ShopID                       string `json:"shopId,omitempty"`
	CreatedAt                    string `json:"createdAt,omitempty"`
	UpdatedAt                    string `json:"updatedAt,omitempty"`
	LastSyncedAt                 string `json:"lastSyncedAt,omitempty"`
	IDCarrier                    string `json:"idCarrier,omitempty"`
	IDReference                  string `json:"idReference,omitempty"`
	CarrierTaxesRatesGroupID     string `json:"carrierTaxesRatesGroupId,omitempty"`
	Name                         string `json:"name,omitempty"`
	URL                          string `json:"url,omitempty"`
	Active                       bool   `json:"active,omitempty"`
	Deleted                      bool   `json:"deleted,omitempty"`
	ShippingHandling             int    `json:"shippingHandling,omitempty"`
	FreeShippingStartsAtPrice    int    `json:"freeShippingStartsAtPrice,omitempty"`
	FreeShippingStartsAtWeight   int    `json:"freeShippingStartsAtWeight,omitempty"`
	DisableCarrierWhenOutOfRange bool   `json:"disableCarrierWhenOutOfRange,omitempty"`
	IsModule                     bool   `json:"isModule,omitempty"`
	IsFree                       bool   `json:"isFree,omitempty"`
	ShippingExternal             bool   `json:"shippingExternal,omitempty"`
	NeedRange                    bool   `json:"needRange,omitempty"`
	ExternalModuleName           string `json:"externalModuleName,omitempty"`
	MaxWidth                     int    `json:"maxWidth,omitempty"`
	MaxHeight                    int    `json:"maxHeight,omitempty"`
	MaxDepth                     int    `json:"maxDepth,omitempty"`
	MaxWeight                    int    `json:"maxWeight,omitempty"`
	Grade                        int    `json:"grade,omitempty"`
	Delay                        string `json:"delay,omitempty"`
	Currency                     string `json:"currency,omitempty"`
	WeightUnit                   string `json:"weightUnit,omitempty"`
	DeletedAt                    string `json:"deletedAt,omitempty"`
}

type CarrierListParams struct {
	Active                       *bool   `url:"active,omitempty"`
	CarrierTaxesRatesGroupId     string  `url:"carrierTaxesRatesGroupId,omitempty"`
	Currency                     string  `url:"currency,omitempty"`
	Delay                        string  `url:"delay,omitempty"`
	Deleted                      *bool   `url:"deleted,omitempty"`
	DeletedAt                    string  `url:"deletedAt,omitempty"`
	DeletedAtStart               string  `url:"deletedAtStart,omitempty"`
	DeletedAtEnd                 string  `url:"deletedAtEnd,omitempty"`
	DisableCarrierWhenOutOfRange *bool   `url:"disableCarrierWhenOutOfRange,omitempty"`
	ExternalModuleName           string  `url:"externalModuleName,omitempty"`
	FreeShippingStartsAtPrice    float64 `url:"freeShippingStartsAtPrice,omitempty"`
	FreeShippingStartsAtPriceStart float64 `url:"freeShippingStartsAtPriceStart,omitempty"`
	FreeShippingStartsAtPriceEnd float64 `url:"freeShippingStartsAtPriceEnd,omitempty"`
	FreeShippingStartsAtWeight   float64 `url:"freeShippingStartsAtWeight,omitempty"`
	FreeShippingStartsAtWeightStart float64 `url:"freeShippingStartsAtWeightStart,omitempty"`
	FreeShippingStartsAtWeightEnd float64 `url:"freeShippingStartsAtWeightEnd,omitempty"`
	Grade                        int     `url:"grade,omitempty"`
	GradeStart                   int     `url:"gradeStart,omitempty"`
	GradeEnd                     int     `url:"gradeEnd,omitempty"`
	IdCarrier                    string  `url:"idCarrier,omitempty"`
	IdReference                  string  `url:"idReference,omitempty"`
	IsFree                       *bool   `url:"isFree,omitempty"`
	IsModule                     *bool   `url:"isModule,omitempty"`
	MaxDepth                     float64 `url:"maxDepth,omitempty"`
	MaxDepthStart                float64 `url:"maxDepthStart,omitempty"`
	MaxDepthEnd                  float64 `url:"maxDepthEnd,omitempty"`
	MaxHeight                    float64 `url:"maxHeight,omitempty"`
	MaxHeightStart               float64 `url:"maxHeightStart,omitempty"`
	MaxHeightEnd                 float64 `url:"maxHeightEnd,omitempty"`
	MaxWeight                    float64 `url:"maxWeight,omitempty"`
	MaxWeightStart               float64 `url:"maxWeightStart,omitempty"`
	MaxWeightEnd                 float64 `url:"maxWeightEnd,omitempty"`
	MaxWidth                     float64 `url:"maxWidth,omitempty"`
	MaxWidthStart                float64 `url:"maxWidthStart,omitempty"`
	MaxWidthEnd                  float64 `url:"maxWidthEnd,omitempty"`
	Name                         string  `url:"name,omitempty"`
	NeedRange                    *bool   `url:"needRange,omitempty"`
	ShippingExternal             *bool   `url:"shippingExternal,omitempty"`
	ShippingHandling             float64 `url:"shippingHandling,omitempty"`
	ShippingHandlingStart        float64 `url:"shippingHandlingStart,omitempty"`
	ShippingHandlingEnd          float64 `url:"shippingHandlingEnd,omitempty"`
	Url                          string  `url:"url,omitempty"`
	WeightUnit                   string  `url:"weightUnit,omitempty"`
	GenericListParams
}

// List carriers. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Carriers_getPaginatedItems
func (service *CarriersService) List(shopID string, opts *CarrierListParams) (*CarriersResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/carriers", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	carriers := new(CarriersResponse)
	response, err := service.client.Do(req, carriers)

	if err != nil {
		return nil, response, err
	}

	return carriers, response, nil
}

// Get a carrier. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Carriers_getSingleItem
func (service *CarriersService) Get(shopID string, shopContentId string) (*CarriersResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/carriers/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	carrier := new(CarriersResponse)
	response, err := service.client.Do(req, carrier)

	if err != nil {
		return nil, response, err
	}

	return carrier, response, nil
}
