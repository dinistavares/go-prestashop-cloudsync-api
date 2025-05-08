package prestashop

import (
	"fmt"
	"net/http"
)

// OrderDetail service
type OrderDetailsService service

type OrderDetailsResponse struct {
	OrderDetail  *OrderDetail   `json:"item,omitempty"`
	OrderDetails *[]OrderDetail `json:"items,omitempty"`
	GenericListResponse
}

type OrderDetail struct {
	ID                 string  `json:"id,omitempty"`
	ShopID             string  `json:"shopId,omitempty"`
	CreatedAt          string  `json:"createdAt,omitempty"`
	UpdatedAt          string  `json:"updatedAt,omitempty"`
	LastSyncedAt       string  `json:"lastSyncedAt,omitempty"`
	IDOrderDetail      string  `json:"idOrderDetail,omitempty"`
	IDOrder            string  `json:"idOrder,omitempty"`
	ProductID          string  `json:"productId,omitempty"`
	ProductAttributeID string  `json:"productAttributeId,omitempty"`
	ProductQuantity    int     `json:"productQuantity,omitempty"`
	UnitPriceTaxIncl   float64 `json:"unitPriceTaxIncl,omitempty"`
	Refund             int     `json:"refund,omitempty"`
	RefundTaxExcl      int     `json:"refundTaxExcl,omitempty"`
	UnitPriceTaxExcl   float64 `json:"unitPriceTaxExcl,omitempty"`
	Currency           string  `json:"currency,omitempty"`
	IsoCode            string  `json:"isoCode,omitempty"`
	UniqueProductID    string  `json:"uniqueProductId,omitempty"`
	Category           string  `json:"category,omitempty"`
	ConversionRate     int     `json:"conversionRate,omitempty"`
}

type OrderDetailListParams struct {
	Category           string  `url:"category,omitempty"`
	ConversionRate     float64 `url:"conversionRate,omitempty"`
	Currency           string  `url:"currency,omitempty"`
	IdOrder            string  `url:"idOrder,omitempty"`
	IdOrderDetail      string  `url:"idOrderDetail,omitempty"`
	IsoCode            string  `url:"isoCode,omitempty"`
	ProductAttributeId float64 `url:"productAttributeId,omitempty"`
	ProductId          float64 `url:"productId,omitempty"`
	ProductQuantity    float64 `url:"productQuantity,omitempty"`
	Refund             float64 `url:"refund,omitempty"`
	RefundTaxExcl      float64 `url:"refundTaxExcl,omitempty"`
	UniqueProductId    string  `url:"uniqueProductId,omitempty"`
	UnitPriceTaxExcl   float64 `url:"unitPriceTaxExcl,omitempty"`
	UnitPriceTaxIncl   float64 `url:"unitPriceTaxIncl,omitempty"`
	GenericListParams
}

// List order details. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/OrderDetails_getPaginatedItems
func (service *OrderDetailsService) List(shopID string, opts *OrderDetailListParams) (*OrderDetailsResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/order-details", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	orderDetails := new(OrderDetailsResponse)
	response, err := service.client.Do(req, orderDetails)

	if err != nil {
		return nil, response, err
	}

	return orderDetails, response, nil
}

// Get a orderDetail. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/OrderDetails_getSingleItem
func (service *OrderDetailsService) Get(shopID string, shopContentId string) (*OrderDetailsResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/order-details/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	orderDetail := new(OrderDetailsResponse)
	response, err := service.client.Do(req, orderDetail)

	if err != nil {
		return nil, response, err
	}

	return orderDetail, response, nil
}
