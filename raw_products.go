package prestashop

import (
	"fmt"
	"net/http"
	"time"
)

// Product service
type ProductsService service

type ProductsResponse struct {
	Product  *Product   `json:"item,omitempty"`
	Products *[]Product `json:"items,omitempty"`
	GenericListResponse
}

type Product struct {
	IDProduct               int        `json:"idProduct,omitempty"`
	IDAttribute             int        `json:"idAttribute,omitempty"`
	IDCategoryDefault       int        `json:"idCategoryDefault,omitempty"`
	Quantity                int        `json:"quantity,omitempty"`
	AdditionalDeliveryTimes int        `json:"additionalDeliveryTimes,omitempty"`
	AdditionalShippingCost  int        `json:"additionalShippingCost,omitempty"`
	UnitPriceRatio          int        `json:"unitPriceRatio,omitempty"`
	IDCategoryDefaultBigint int        `json:"idCategoryDefaultBigint,omitempty"`
	IDManufacturer          int        `json:"idManufacturer,omitempty"`
	IDSupplier              int        `json:"idSupplier,omitempty"`
	Weight                  float64    `json:"weight,omitempty"`
	PriceTaxIncl            float64    `json:"priceTaxIncl,omitempty"`
	PriceTaxExcl            float64    `json:"priceTaxExcl,omitempty"`
	SalePriceTaxExcl        float64    `json:"salePriceTaxExcl,omitempty"`
	SalePriceTaxIncl        float64    `json:"salePriceTaxIncl,omitempty"`
	Width                   float64    `json:"width,omitempty"`
	Height                  float64    `json:"height,omitempty"`
	Depth                   float64    `json:"depth,omitempty"`
	Tax                     float64    `json:"tax,omitempty"`
	SaleTax                 float64    `json:"saleTax,omitempty"`
	ID                      string     `json:"id,omitempty"`
	ShopID                  string     `json:"shopId,omitempty"`
	UniqueProductID         string     `json:"uniqueProductId,omitempty"`
	Name                    string     `json:"name,omitempty"`
	Description             string     `json:"description,omitempty"`
	DescriptionShort        string     `json:"descriptionShort,omitempty"`
	IsoCode                 string     `json:"isoCode,omitempty"`
	DefaultCategory         string     `json:"defaultCategory,omitempty"`
	LinkRewrite             string     `json:"linkRewrite,omitempty"`
	Reference               string     `json:"reference,omitempty"`
	Upc                     string     `json:"upc,omitempty"`
	Ean                     string     `json:"ean,omitempty"`
	Isbn                    string     `json:"isbn,omitempty"`
	Condition               string     `json:"condition,omitempty"`
	Visibility              string     `json:"visibility,omitempty"`
	Manufacturer            string     `json:"manufacturer,omitempty"`
	Link                    string     `json:"link,omitempty"`
	Cover                   string     `json:"cover,omitempty"`
	Attributes              string     `json:"attributes,omitempty"`
	Features                string     `json:"features,omitempty"`
	Images                  string     `json:"images,omitempty"`
	CategoryPath            string     `json:"categoryPath,omitempty"`
	CategoryIDPath          string     `json:"categoryIdPath,omitempty"`
	IDProductAttribute      string     `json:"idProductAttribute,omitempty"`
	AvailableDateLegacy     string     `json:"availableDateLegacy,omitempty"`
	DeliveryInStock         string     `json:"deliveryInStock,omitempty"`
	DeliveryOutStock        string     `json:"deliveryOutStock,omitempty"`
	SaleDate                string     `json:"saleDate,omitempty"`
	Unity                   string     `json:"unity,omitempty"`
	PricePerUnit            string     `json:"pricePerUnit,omitempty"`
	Mpn                     string     `json:"mpn,omitempty"`
	Active                  *bool      `json:"active,omitempty"`
	Deleted                 *bool      `json:"deleted,omitempty"`
	IsDefaultAttribute      *bool      `json:"isDefaultAttribute,omitempty"`
	AvailableForOrder       *bool      `json:"availableForOrder,omitempty"`
	IsBundle                *bool      `json:"isBundle,omitempty"`
	IsVirtual               *bool      `json:"isVirtual,omitempty"`
	CreatedAt               *time.Time `json:"createdAt,omitempty"`
	UpdatedAt               *time.Time `json:"updatedAt,omitempty"`
	LastSyncedAt            *time.Time `json:"lastSyncedAt,omitempty"`
	DeletedAt               *time.Time `json:"deletedAt,omitempty"`
	AvailableDate           *time.Time `json:"availableDate,omitempty"`
}

type ProductListParams struct {
	IdAttribute                  int     `url:"idAttribute,omitempty"`
	IdAttributeStart             int     `url:"idAttributeStart,omitempty"`
	IdAttributeEnd               int     `url:"idAttributeEnd,omitempty"`
	IdCategoryDefault            int     `url:"idCategoryDefault,omitempty"`
	IdCategoryDefaultStart       int     `url:"idCategoryDefaultStart,omitempty"`
	IdCategoryDefaultEnd         int     `url:"idCategoryDefaultEnd,omitempty"`
	IdCategoryDefaultBigint      int     `url:"idCategoryDefaultBigint,omitempty"`
	IdCategoryDefaultBigintStart int     `url:"idCategoryDefaultBigintStart,omitempty"`
	IdCategoryDefaultBigintEnd   int     `url:"idCategoryDefaultBigintEnd,omitempty"`
	IdManufacturer               int     `url:"idManufacturer,omitempty"`
	IdProduct                    int     `url:"idProduct,omitempty"`
	IdProductStart               int     `url:"idProductStart,omitempty"`
	IdProductEnd                 int     `url:"idProductEnd,omitempty"`
	IdSupplier                   int     `url:"idSupplier,omitempty"`
	Quantity                     int     `url:"quantity,omitempty"`
	QuantityStart                int     `url:"quantityStart,omitempty"`
	QuantityEnd                  int     `url:"quantityEnd,omitempty"`
	AdditionalDeliveryTimes      float64 `url:"additionalDeliveryTimes,omitempty"`
	AdditionalDeliveryTimesStart float64 `url:"additionalDeliveryTimesStart,omitempty"`
	AdditionalDeliveryTimesEnd   float64 `url:"additionalDeliveryTimesEnd,omitempty"`
	AdditionalShippingCost       float64 `url:"additionalShippingCost,omitempty"`
	AdditionalShippingCostStart  float64 `url:"additionalShippingCostStart,omitempty"`
	AdditionalShippingCostEnd    float64 `url:"additionalShippingCostEnd,omitempty"`
	Depth                        float64 `url:"depth,omitempty"`
	DepthStart                   float64 `url:"depthStart,omitempty"`
	DepthEnd                     float64 `url:"depthEnd,omitempty"`
	Height                       float64 `url:"height,omitempty"`
	HeightStart                  float64 `url:"heightStart,omitempty"`
	HeightEnd                    float64 `url:"heightEnd,omitempty"`
	PricePerUnit                 float64 `url:"pricePerUnit,omitempty"`
	PricePerUnitStart            float64 `url:"pricePerUnitStart,omitempty"`
	PricePerUnitEnd              float64 `url:"pricePerUnitEnd,omitempty"`
	PriceTaxExcl                 float64 `url:"priceTaxExcl,omitempty"`
	PriceTaxExclStart            float64 `url:"priceTaxExclStart,omitempty"`
	PriceTaxExclEnd              float64 `url:"priceTaxExclEnd,omitempty"`
	PriceTaxIncl                 float64 `url:"priceTaxIncl,omitempty"`
	PriceTaxInclStart            float64 `url:"priceTaxInclStart,omitempty"`
	PriceTaxInclEnd              float64 `url:"priceTaxInclEnd,omitempty"`
	SalePriceTaxExcl             float64 `url:"salePriceTaxExcl,omitempty"`
	SalePriceTaxExclStart        float64 `url:"salePriceTaxExclStart,omitempty"`
	SalePriceTaxExclEnd          float64 `url:"salePriceTaxExclEnd,omitempty"`
	SalePriceTaxIncl             float64 `url:"salePriceTaxIncl,omitempty"`
	SalePriceTaxInclStart        float64 `url:"salePriceTaxInclStart,omitempty"`
	SalePriceTaxInclEnd          float64 `url:"salePriceTaxInclEnd,omitempty"`
	SaleTax                      float64 `url:"saleTax,omitempty"`
	SaleTaxStart                 float64 `url:"saleTaxStart,omitempty"`
	SaleTaxEnd                   float64 `url:"saleTaxEnd,omitempty"`
	Tax                          float64 `url:"tax,omitempty"`
	TaxStart                     float64 `url:"taxStart,omitempty"`
	TaxEnd                       float64 `url:"taxEnd,omitempty"`
	UnitPriceRatio               float64 `url:"unitPriceRatio,omitempty"`
	UnitPriceRatioStart          float64 `url:"unitPriceRatioStart,omitempty"`
	UnitPriceRatioEnd            float64 `url:"unitPriceRatioEnd,omitempty"`
	Weight                       float64 `url:"weight,omitempty"`
	WeightStart                  float64 `url:"weightStart,omitempty"`
	WeightEnd                    float64 `url:"weightEnd,omitempty"`
	Width                        float64 `url:"width,omitempty"`
	WidthStart                   float64 `url:"widthStart,omitempty"`
	WidthEnd                     float64 `url:"widthEnd,omitempty"`
	Active                       *bool   `url:"active,omitempty"`
	AvailableForOrder            *bool   `url:"availableForOrder,omitempty"`
	Deleted                      *bool   `url:"deleted,omitempty"`
	IsBundle                     *bool   `url:"isBundle,omitempty"`
	IsDefaultAttribute           *bool   `url:"isDefaultAttribute,omitempty"`
	IsVirtual                    *bool   `url:"isVirtual,omitempty"`
	DeletedAtStart               string  `url:"deletedAtStart,omitempty"`
	DeletedAtEnd                 string  `url:"deletedAtEnd,omitempty"`
	DeliveryInStock              string  `url:"deliveryInStock,omitempty"`
	DeliveryOutStock             string  `url:"deliveryOutStock,omitempty"`
	Attributes                   string  `url:"attributes,omitempty"`
	AvailableDate                string  `url:"availableDate,omitempty"`
	AvailableDateStart           string  `url:"availableDateStart,omitempty"`
	AvailableDateEnd             string  `url:"availableDateEnd,omitempty"`
	AvailableDateLegacy          string  `url:"availableDateLegacy,omitempty"`
	CategoryIdPath               string  `url:"categoryIdPath,omitempty"`
	CategoryPath                 string  `url:"categoryPath,omitempty"`
	Condition                    string  `url:"condition,omitempty"`
	Cover                        string  `url:"cover,omitempty"`
	DefaultCategory              string  `url:"defaultCategory,omitempty"`
	DeletedAt                    string  `url:"deletedAt,omitempty"`
	Description                  string  `url:"description,omitempty"`
	DescriptionShort             string  `url:"descriptionShort,omitempty"`
	Ean                          string  `url:"ean,omitempty"`
	Features                     string  `url:"features,omitempty"`
	IdProductAttribute           string  `url:"idProductAttribute,omitempty"`
	Images                       string  `url:"images,omitempty"`
	Isbn                         string  `url:"isbn,omitempty"`
	IsoCode                      string  `url:"isoCode,omitempty"`
	Link                         string  `url:"link,omitempty"`
	LinkRewrite                  string  `url:"linkRewrite,omitempty"`
	Manufacturer                 string  `url:"manufacturer,omitempty"`
	Mpn                          string  `url:"mpn,omitempty"`
	Name                         string  `url:"name,omitempty"`
	Reference                    string  `url:"reference,omitempty"`
	SaleDate                     string  `url:"saleDate,omitempty"`
	UniqueProductId              string  `url:"uniqueProductId,omitempty"`
	Unity                        string  `url:"unity,omitempty"`
	Upc                          string  `url:"upc,omitempty"`
	Visibility                   string  `url:"visibility,omitempty"`
	GenericListParams
}

// List products. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Products_getPaginatedItems
func (service *ProductsService) List(shopID string, opts *ProductListParams) (*ProductsResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/products", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	products := new(ProductsResponse)
	response, err := service.client.Do(req, products)

	if err != nil {
		return nil, response, err
	}

	return products, response, nil
}

// Get a product. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Products_getSingleItem
func (service *ProductsService) Get(shopID string, shopContentId string) (*ProductsResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/products/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	product := new(ProductsResponse)
	response, err := service.client.Do(req, product)

	if err != nil {
		return nil, response, err
	}

	return product, response, nil
}
