package prestashop

import (
	"fmt"
	"net/http"
)

// Currency service
type CurrenciesService service

type CurrenciesResponse struct {
	Currency   Currency   `json:"item,omitempty"`
	Currencies []Currency `json:"items,omitempty"`
	GenericListResponse
}

type Currency struct {
	ID               string `json:"id,omitempty"`
	ShopID           string `json:"shopId,omitempty"`
	CreatedAt        string `json:"createdAt,omitempty"`
	UpdatedAt        string `json:"updatedAt,omitempty"`
	LastSyncedAt     string `json:"lastSyncedAt,omitempty"`
	IDCurrency       int    `json:"idCurrency,omitempty"`
	IDParent         int    `json:"idParent,omitempty"`
	UniqueCurrencyID string `json:"uniqueCurrencyId,omitempty"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	LinkRewrite      string `json:"linkRewrite,omitempty"`
	MetaTitle        string `json:"metaTitle,omitempty"`
	MetaKeywords     string `json:"metaKeywords,omitempty"`
	MetaDescription  string `json:"metaDescription,omitempty"`
	IsoCode          string `json:"isoCode,omitempty"`
	DeletedAt        string `json:"deletedAt,omitempty"`
	Deleted          bool   `json:"deleted,omitempty"`
}

type CurrencyListParams struct {
	IdCurrency       int    `url:"idCurrency,omitempty"`
	IdCurrencyStart  int    `url:"idCurrencyStart,omitempty"`
	IdCurrencyEnd    int    `url:"idCurrencyEnd,omitempty"`
	IdParent         int    `url:"idParent,omitempty"`
	IdParentStart    int    `url:"idParentStart,omitempty"`
	IdParentEnd      int    `url:"idParentEnd,omitempty"`
	IsoCode          string `url:"isoCode,omitempty"`
	LinkRewrite      string `url:"linkRewrite,omitempty"`
	MetaDescription  string `url:"metaDescription,omitempty"`
	MetaKeywords     string `url:"metaKeywords,omitempty"`
	MetaTitle        string `url:"metaTitle,omitempty"`
	Name             string `url:"name,omitempty"`
	UniqueCurrencyId string `url:"uniqueCurrencyId,omitempty"`
	Deleted          *bool  `url:"deleted,omitempty"`
	Description      string `url:"description,omitempty"`
	GenericListParams
}

// List currencies. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Currencies_getPaginatedItems
func (service *CurrenciesService) List(shopID string, opts *CurrencyListParams) (*CurrenciesResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/currencies", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	currencies := new(CurrenciesResponse)
	response, err := service.client.Do(req, currencies)

	if err != nil {
		return nil, response, err
	}

	return currencies, response, nil
}

// Get a category. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Currencies_getSingleItem
func (service *CurrenciesService) Get(shopID string, shopContentId string) (*CurrenciesResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/currencies/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	category := new(CurrenciesResponse)
	response, err := service.client.Do(req, category)

	if err != nil {
		return nil, response, err
	}

	return category, response, nil
}
