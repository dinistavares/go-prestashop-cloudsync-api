package prestashop

import (
	"fmt"
	"net/http"
)

// OrderStatusHistory service
type OrderStatusHistoryService service

type OrderStatusHistoryResponse struct {
	OrderStatusHistory   *OrderStatusHistory   `json:"item,omitempty"`
	OrderStatusHistories *[]OrderStatusHistory `json:"items,omitempty"`
	GenericListResponse
}

type OrderStatusHistory struct {
	ID               string   `json:"id,omitempty"`
	ShopID           string   `json:"shopId,omitempty"`
	CreatedAt        string   `json:"createdAt,omitempty"`
	UpdatedAt        string   `json:"updatedAt,omitempty"`
	LastSyncedAt     string   `json:"lastSyncedAt,omitempty"`
	IDOrderHistory   *int     `json:"idOrderHistory,omitempty"`
	IDOrder          string   `json:"idOrder,omitempty"`
	DateAdd          string   `json:"dateAdd,omitempty"`
	IDOrderState     *int     `json:"idOrderState,omitempty"`
	Name             string   `json:"name,omitempty"`
	Template         string   `json:"template,omitempty"`
	IsValidated      *bool    `json:"isValidated,omitempty"`
	IsDelivered      *bool    `json:"isDelivered,omitempty"`
	IsShipped        *bool    `json:"isShipped,omitempty"`
	IsPaid           *bool    `json:"isPaid,omitempty"`
	IsDeleted        *bool    `json:"isDeleted,omitempty"`
	TotalPaidTaxExcl *float64 `json:"totalPaidTaxExcl,omitempty"`
	RefundTaxExcl    *float64 `json:"refundTaxExcl,omitempty"`
}

type OrderStatusHistoryListParams struct {
	DateAdd          string `url:"dateAdd,omitempty"`
	IdOrder          string `url:"idOrder,omitempty"`
	IdOrderHistory   int    `url:"idOrderHistory,omitempty"`
	IdOrderState     string `url:"idOrderState,omitempty"`
	IsDeleted        bool   `url:"isDeleted,omitempty"`
	IsDelivered      bool   `url:"isDelivered,omitempty"`
	IsPaid           bool   `url:"isPaid,omitempty"`
	IsShipped        bool   `url:"isShipped,omitempty"`
	IsValidated      bool   `url:"isValidated,omitempty"`
	Name             string `url:"name,omitempty"`
	RefundTaxExcl    int    `url:"refundTaxExcl,omitempty"`
	Template         string `url:"template,omitempty"`
	TotalPaidTaxExcl int    `url:"totalPaidTaxExcl,omitempty"`
	GenericListParams
}

// List order details. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/OrderStatusHistory_getPaginatedItems
func (service *OrderStatusHistoryService) List(shopID string, opts *OrderStatusHistoryListParams) (*OrderStatusHistoryResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/order-status-history", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	orderDetails := new(OrderStatusHistoryResponse)
	response, err := service.client.Do(req, orderDetails)

	if err != nil {
		return nil, response, err
	}

	return orderDetails, response, nil
}

// Get a orderDetail. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/OrderStatusHistory_getSingleItem
func (service *OrderStatusHistoryService) Get(shopID string, shopContentId string) (*OrderStatusHistoryResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/order-status-history/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	orderDetail := new(OrderStatusHistoryResponse)
	response, err := service.client.Do(req, orderDetail)

	if err != nil {
		return nil, response, err
	}

	return orderDetail, response, nil
}
