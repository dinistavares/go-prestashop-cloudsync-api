package prestashop

import (
	"fmt"
	"net/http"
	"time"
)

// Customer service
type OrdersService service

type OrdersResponse struct {
	Customer  *Order   `json:"item,omitempty"`
	Customers *[]Order `json:"items,omitempty"`
	GenericListResponse
}

type Order struct {
	CurrentState        float64    `json:"currentState,omitempty"`
	ConversionRate      float64    `json:"conversionRate,omitempty"`
	TotalPaidTaxExcl    float64    `json:"totalPaidTaxExcl,omitempty"`
	Refund              float64    `json:"refund,omitempty"`
	RefundTaxExcl       float64    `json:"refundTaxExcl,omitempty"`
	ShippingCost        float64    `json:"shippingCost,omitempty"`
	TotalPaidTax        float64    `json:"totalPaidTax,omitempty"`
	TotalPaidReal       float64    `json:"totalPaidReal,omitempty"`
	IDCarrier           float64    `json:"idCarrier,omitempty"`
	TotalPaidTaxIncl    float64    `json:"totalPaidTaxIncl,omitempty"`
	ID                  string     `json:"id,omitempty"`
	ShopID              string     `json:"shopId,omitempty"`
	IDOrder             string     `json:"idOrder,omitempty"`
	Reference           string     `json:"reference,omitempty"`
	IDCustomer          string     `json:"idCustomer,omitempty"`
	IDCart              string     `json:"idCart,omitempty"`
	Currency            string     `json:"currency,omitempty"`
	PaymentMode         string     `json:"paymentMode,omitempty"`
	PaymentModule       string     `json:"paymentModule,omitempty"`
	InvoiceCountryCode  string     `json:"invoiceCountryCode,omitempty"`
	DeliveryCountryCode string     `json:"deliveryCountryCode,omitempty"`
	NewCustomer         *bool      `json:"newCustomer,omitempty"`
	IsPaid              *bool      `json:"isPaid,omitempty"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
	CreatedAt           *time.Time `json:"createdAt,omitempty"`
	LastSyncedAt        *time.Time `json:"lastSyncedAt,omitempty"`
}

type OrderListParams struct {
	ConversionRate        float64 `url:"conversionRate,omitempty"`
	ConversionRateStart   float64 `url:"conversionRateStart,omitempty"`
	ConversionRateEnd     float64 `url:"conversionRateEnd,omitempty"`
	Currency              string  `url:"currency,omitempty"`
	CurrentState          int     `url:"currentState,omitempty"`
	CurrentStateStart     int     `url:"currentStateStart,omitempty"`
	CurrentStateEnd       int     `url:"currentStateEnd,omitempty"`
	DeliveryCountryCode   string  `url:"deliveryCountryCode,omitempty"`
	Direction             string  `url:"direction,omitempty"`
	IdCarrier             int     `url:"idCarrier,omitempty"`
	IdCarrierStart        int     `url:"idCarrierStart,omitempty"`
	IdCarrierEnd          int     `url:"idCarrierEnd,omitempty"`
	IdCart                string  `url:"idCart,omitempty"`
	IdCustomer            string  `url:"idCustomer,omitempty"`
	IdOrder               string  `url:"idOrder,omitempty"`
	InvoiceCountryCode    string  `url:"invoiceCountryCode,omitempty"`
	IsPaid                *bool   `url:"isPaid,omitempty"`
	NewCustomer           *bool   `url:"newCustomer,omitempty"`
	PaymentMode           string  `url:"paymentMode,omitempty"`
	PaymentModule         string  `url:"paymentModule,omitempty"`
	Reference             string  `url:"reference,omitempty"`
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
	GenericListParams
}

// List customers. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Orders_getPaginatedItems
func (service *OrdersService) List(shopID string, opts *OrderListParams) (*OrdersResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/orders", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	customers := new(OrdersResponse)
	response, err := service.client.Do(req, customers)

	if err != nil {
		return nil, response, err
	}

	return customers, response, nil
}

// Get a customer. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Orders_getSingleItem
func (service *OrdersService) Get(shopID string, shopContentId string) (*OrdersResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/orders/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	customer := new(OrdersResponse)
	response, err := service.client.Do(req, customer)

	if err != nil {
		return nil, response, err
	}

	return customer, response, nil
}
