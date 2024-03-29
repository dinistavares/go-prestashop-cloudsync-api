package prestashop

import (
	"fmt"
	"net/http"
)

// Order service
type OrdersService service

type OrdersResponse struct {
	Order  *Order   `json:"item,omitempty"`
	Orders *[]Order `json:"items,omitempty"`
	GenericListResponse
}

type Order struct {
	IDCarrier           int     `json:"idCarrier,omitempty"`
	CurrentState        int     `json:"currentState,omitempty"`
	ConversionRate      float64 `json:"conversionRate,omitempty"`
	TotalPaidTaxExcl    float64 `json:"totalPaidTaxExcl,omitempty"`
	Refund              float64 `json:"refund,omitempty"`
	RefundTaxExcl       float64 `json:"refundTaxExcl,omitempty"`
	ShippingCost        float64 `json:"shippingCost,omitempty"`
	TotalPaidTax        float64 `json:"totalPaidTax,omitempty"`
	TotalPaidReal       float64 `json:"totalPaidReal,omitempty"`
	TotalPaidTaxIncl    float64 `json:"totalPaidTaxIncl,omitempty"`
	ID                  string  `json:"id,omitempty"`
	ShopID              string  `json:"shopId,omitempty"`
	IDOrder             string  `json:"idOrder,omitempty"`
	Reference           string  `json:"reference,omitempty"`
	IDCustomer          string  `json:"idCustomer,omitempty"`
	IDCart              string  `json:"idCart,omitempty"`
	Currency            string  `json:"currency,omitempty"`
	PaymentMode         string  `json:"paymentMode,omitempty"`
	PaymentModule       string  `json:"paymentModule,omitempty"`
	InvoiceCountryCode  string  `json:"invoiceCountryCode,omitempty"`
	DeliveryCountryCode string  `json:"deliveryCountryCode,omitempty"`
	NewCustomer         *bool   `json:"newCustomer,omitempty"`
	IsPaid              *bool   `json:"isPaid,omitempty"`
	UpdatedAt           string  `json:"updatedAt,omitempty"`
	CreatedAt           string  `json:"createdAt,omitempty"`
	LastSyncedAt        string  `json:"lastSyncedAt,omitempty"`
}

type OrderListParams struct {
	IdCarrier             int     `url:"idCarrier,omitempty"`
	CurrentState          int     `url:"currentState,omitempty"`
	CurrentStateStart     int     `url:"currentStateStart,omitempty"`
	CurrentStateEnd       int     `url:"currentStateEnd,omitempty"`
	ConversionRate        float64 `url:"conversionRate,omitempty"`
	ConversionRateStart   float64 `url:"conversionRateStart,omitempty"`
	ConversionRateEnd     float64 `url:"conversionRateEnd,omitempty"`
	Refund                float64 `url:"refund,omitempty"`
	RefundStart           float64 `url:"refundStart,omitempty"`
	RefundEnd             float64 `url:"refundEnd,omitempty"`
	RefundTaxExcl         float64 `url:"refundTaxExcl,omitempty"`
	RefundTaxExclStart    float64 `url:"refundTaxExclStart,omitempty"`
	RefundTaxExclEnd      float64 `url:"refundTaxExclEnd,omitempty"`
	ShippingCost          float64 `url:"shippingCost,omitempty"`
	ShippingCostStart     float64 `url:"shippingCostStart,omitempty"`
	ShippingCostEnd       float64 `url:"shippingCostEnd,omitempty"`
	TotalPaidReal         float64 `url:"totalPaidReal,omitempty"`
	TotalPaidRealStart    float64 `url:"totalPaidRealStart,omitempty"`
	TotalPaidRealEnd      float64 `url:"totalPaidRealEnd,omitempty"`
	TotalPaidTax          float64 `url:"totalPaidTax,omitempty"`
	TotalPaidTaxStart     float64 `url:"totalPaidTaxStart,omitempty"`
	TotalPaidTaxEnd       float64 `url:"totalPaidTaxEnd,omitempty"`
	TotalPaidTaxExcl      float64 `url:"totalPaidTaxExcl,omitempty"`
	TotalPaidTaxExclStart float64 `url:"totalPaidTaxExclStart,omitempty"`
	TotalPaidTaxExclEnd   float64 `url:"totalPaidTaxExclEnd,omitempty"`
	TotalPaidTaxIncl      float64 `url:"totalPaidTaxIncl,omitempty"`
	TotalPaidTaxInclStart float64 `url:"totalPaidTaxInclStart,omitempty"`
	TotalPaidTaxInclEnd   float64 `url:"totalPaidTaxInclEnd,omitempty"`
	Currency              string  `url:"currency,omitempty"`
	IdCart                string  `url:"idCart,omitempty"`
	IdCustomer            string  `url:"idCustomer,omitempty"`
	IdOrder               string  `url:"idOrder,omitempty"`
	DeliveryCountryCode   string  `url:"deliveryCountryCode,omitempty"`
	InvoiceCountryCode    string  `url:"invoiceCountryCode,omitempty"`
	PaymentMode           string  `url:"paymentMode,omitempty"`
	PaymentModule         string  `url:"paymentModule,omitempty"`
	Reference             string  `url:"reference,omitempty"`
	IsPaid                *bool   `url:"isPaid,omitempty"`
	NewCustomer           *bool   `url:"newCustomer,omitempty"`
	GenericListParams
}

// List orders. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Orders_getPaginatedItems
func (service *OrdersService) List(shopID string, opts *OrderListParams) (*OrdersResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/orders", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	orders := new(OrdersResponse)
	response, err := service.client.Do(req, orders)

	if err != nil {
		return nil, response, err
	}

	return orders, response, nil
}

// Get a order. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Orders_getSingleItem
func (service *OrdersService) Get(shopID string, shopContentId string) (*OrdersResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/orders/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	order := new(OrdersResponse)
	response, err := service.client.Do(req, order)

	if err != nil {
		return nil, response, err
	}

	return order, response, nil
}
