package prestashop

import (
	"fmt"
	"net/http"
)

type ShopResponse struct {
	Kind string `json:"kind,omitempty"`
	Item Item   `json:"item,omitempty"`
}

type Item struct {
	ID                    string `json:"id,omitempty"`
	ShopID                string `json:"shopId,omitempty"`
	CreatedAt             string `json:"createdAt,omitempty"`
	UpdatedAt             string `json:"updatedAt,omitempty"`
	LastSyncedAt          string `json:"lastSyncedAt,omitempty"`
	CmsVersion            string `json:"cmsVersion,omitempty"`
	URLIsSimplified       bool   `json:"urlIsSimplified,omitempty"`
	CartIsPersistent      bool   `json:"cartIsPersistent,omitempty"`
	DefaultLanguage       string `json:"defaultLanguage,omitempty"`
	Languages             string `json:"languages,omitempty"`
	DefaultCurrency       string `json:"defaultCurrency,omitempty"`
	Currencies            string `json:"currencies,omitempty"`
	Timezone              string `json:"timezone,omitempty"`
	PhpVersion            string `json:"phpVersion,omitempty"`
	HTTPServer            string `json:"httpServer,omitempty"`
	IsOrderReturnEnabled  bool   `json:"isOrderReturnEnabled,omitempty"`
	Ssl                   bool   `json:"ssl,omitempty"`
	OrderReturnNbDays     int    `json:"orderReturnNbDays,omitempty"`
	URL                   string `json:"url,omitempty"`
	WeightUnit            string `json:"weightUnit,omitempty"`
	LanguageChangedCount  int    `json:"languageChangedCount,omitempty"`
	LastLanguageChangedAt string `json:"lastLanguageChangedAt,omitempty"`
	MultiShopCount        int    `json:"multiShopCount,omitempty"`
	DimensionUnit         string `json:"dimensionUnit,omitempty"`
	DistanceUnit          string `json:"distanceUnit,omitempty"`
	VolumeUnit            string `json:"volumeUnit,omitempty"`
	CountryCode           string `json:"countryCode,omitempty"`
	FolderCreatedAt       string `json:"folderCreatedAt,omitempty"`
}

// Shop service
type ShopService service

// Get a shop info. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/ShopsController_getSingleItem
func (service *ShopService) Get(shopID string) (*ShopResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/shop", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	shop := new(ShopResponse)
	response, err := service.client.Do(req, shop)

	if err != nil {
		return nil, response, err
	}

	return shop, response, nil
}
