package prestashop

import (
	"fmt"
	"net/http"
)

// Bundle service
type BundlesService service

type BundlesResponse struct {
	Bundle  *Bundle   `json:"item,omitempty"`
	Bundles *[]Bundle `json:"items,omitempty"`
	GenericListResponse
}

type Bundle struct {
	ID                 string `json:"id,omitempty"`
	ShopID             string `json:"shopId,omitempty"`
	CreatedAt          string `json:"createdAt,omitempty"`
	UpdatedAt          string `json:"updatedAt,omitempty"`
	LastSyncedAt       string `json:"lastSyncedAt,omitempty"`
	IDBundleProduct    string `json:"idBundleProduct,omitempty"`
	IDBundle           int    `json:"idBundle,omitempty"`
	IDProduct          int    `json:"idProduct,omitempty"`
	IDProductAttribute int    `json:"idProductAttribute,omitempty"`
	UniqueProductID    string `json:"uniqueProductId,omitempty"`
	Quantity           int    `json:"quantity,omitempty"`
}

type BundleListParams struct {
	IdBundle                int     `url:"idBundle,omitempty"`
	IdBundleStart           int     `url:"idBundleStart,omitempty"`
	IdBundleEnd             int     `url:"idBundleEnd,omitempty"`
	IdProduct               int     `url:"idProduct,omitempty"`
	IdProductStart          int     `url:"idProductStart,omitempty"`
	IdProductEnd            int     `url:"idProductEnd,omitempty"`
	IdProductAttribute      int     `url:"idProductAttribute,omitempty"`
	IdProductAttributeStart int     `url:"idProductAttributeStart,omitempty"`
	IdProductAttributeEnd   int     `url:"idProductAttributeEnd,omitempty"`
	Quantity                int     `url:"quantity,omitempty"`
	QuantityStart           int     `url:"quantityStart,omitempty"`
	QuantityEnd             int     `url:"quantityEnd,omitempty"`
	IdBundleProduct         string  `url:"idBundleProduct,omitempty"`
	UniqueProductId         string  `url:"uniqueProductId,omitempty"`
	GenericListParams
}

// List bundles. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Bundles_getPaginatedItems
func (service *BundlesService) List(shopID string, opts *BundleListParams) (*BundlesResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/bundles", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	bundles := new(BundlesResponse)
	response, err := service.client.Do(req, bundles)

	if err != nil {
		return nil, response, err
	}

	return bundles, response, nil
}

// Get a bundle. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Bundles_getSingleItem
func (service *BundlesService) Get(shopID string, shopContentId string) (*BundlesResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/bundles/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	bundle := new(BundlesResponse)
	response, err := service.client.Do(req, bundle)

	if err != nil {
		return nil, response, err
	}

	return bundle, response, nil
}
