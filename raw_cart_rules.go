package prestashop

import (
	"fmt"
	"net/http"
)

// Cart Rules service
type CartRulesService service

type CartRulesResponse struct {
	CartRule  *CartRule   `json:"item,omitempty"`
	CartRules *[]CartRule `json:"items,omitempty"`
	GenericListResponse
}

type CartRule struct {
	ID                      string  `json:"id,omitempty"`
	ShopID                  string  `json:"shopId,omitempty"`
	CreatedAt               string  `json:"createdAt,omitempty"`
	UpdatedAt               string  `json:"updatedAt,omitempty"`
	LastSyncedAt            string  `json:"lastSyncedAt,omitempty"`
	IDCartRule              int     `json:"idCartRule,omitempty"`
	IDCustomer              int     `json:"idCustomer,omitempty"`
	Code                    string  `json:"code,omitempty"`
	From                    string  `json:"from,omitempty"`
	To                      string  `json:"to,omitempty"`
	Description             string  `json:"description,omitempty"`
	Quantity                int     `json:"quantity,omitempty"`
	QuantityPerUser         int     `json:"quantityPerUser,omitempty"`
	Priority                int     `json:"priority,omitempty"`
	PartialUse              *bool   `json:"partialUse,omitempty"`
	MinimumAmount           float64 `json:"minimumAmount,omitempty"`
	MinimumAmountTax        *bool   `json:"minimumAmountTax,omitempty"`
	MinimumAmountCurrency   int     `json:"minimumAmountCurrency,omitempty"`
	MinimumAmountShipping   *bool   `json:"minimumAmountShipping,omitempty"`
	CountryRestriction      *bool   `json:"countryRestriction,omitempty"`
	CarrierRestriction      *bool   `json:"carrierRestriction,omitempty"`
	GroupRestriction        *bool   `json:"groupRestriction,omitempty"`
	CartRuleRestriction     *bool   `json:"cartRuleRestriction,omitempty"`
	ProductRestriction      *bool   `json:"productRestriction,omitempty"`
	ShopRestriction         *bool   `json:"shopRestriction,omitempty"`
	FreeShipping            *bool   `json:"freeShipping,omitempty"`
	ReductionPercent        float64 `json:"reductionPercent,omitempty"`
	ReductionAmount         float64 `json:"reductionAmount,omitempty"`
	ReductionTax            *bool   `json:"reductionTax,omitempty"`
	ReductionCurrency       int     `json:"reductionCurrency,omitempty"`
	ReductionProduct        int     `json:"reductionProduct,omitempty"`
	ReductionExcludeSpecial *bool   `json:"reductionExcludeSpecial,omitempty"`
	GiftProduct             int     `json:"giftProduct,omitempty"`
	GiftProductAttribute    int     `json:"giftProductAttribute,omitempty"`
	Highlight               *bool   `json:"highlight,omitempty"`
	Active                  *bool   `json:"active,omitempty"`
}

type CartRuleListParams struct {
	Active                     *bool   `url:"active,omitempty"`
	CarrierRestriction         *bool   `url:"carrierRestriction,omitempty"`
	CartRuleRestriction        *bool   `url:"cartRuleRestriction,omitempty"`
	Code                       string  `url:"code,omitempty"`
	CountryRestriction         *bool   `url:"countryRestriction,omitempty"`
	Description                string  `url:"description,omitempty"`
	FreeShipping               *bool   `url:"freeShipping,omitempty"`
	From                       string  `url:"from,omitempty"`
	FromStart                  string  `url:"fromStart,omitempty"`
	FromEnd                    string  `url:"fromEnd,omitempty"`
	GiftProduct                int     `url:"giftProduct,omitempty"`
	GiftProductStart           int     `url:"giftProductStart,omitempty"`
	GiftProductEnd             int     `url:"giftProductEnd,omitempty"`
	GiftProductAttribute       int     `url:"giftProductAttribute,omitempty"`
	GiftProductAttributeStart  int     `url:"giftProductAttributeStart,omitempty"`
	GiftProductAttributeEnd    int     `url:"giftProductAttributeEnd,omitempty"`
	GroupRestriction           *bool   `url:"groupRestriction,omitempty"`
	Highlight                  *bool   `url:"highlight,omitempty"`
	IdCartRule                 int     `url:"idCartRule,omitempty"`
	IdCartRuleStart            int     `url:"idCartRuleStart,omitempty"`
	IdCartRuleEnd              int     `url:"idCartRuleEnd,omitempty"`
	IdCustomer                 int     `url:"idCustomer,omitempty"`
	IdCustomerStart            int     `url:"idCustomerStart,omitempty"`
	IdCustomerEnd              int     `url:"idCustomerEnd,omitempty"`
	MinimumAmount              float64 `url:"minimumAmount,omitempty"`
	MinimumAmountStart         float64 `url:"minimumAmountStart,omitempty"`
	MinimumAmountEnd           float64 `url:"minimumAmountEnd,omitempty"`
	MinimumAmountCurrency      int     `url:"minimumAmountCurrency,omitempty"`
	MinimumAmountCurrencyStart int     `url:"minimumAmountCurrencyStart,omitempty"`
	MinimumAmountCurrencyEnd   int     `url:"minimumAmountCurrencyEnd,omitempty"`
	MinimumAmountShipping      *bool   `url:"minimumAmountShipping,omitempty"`
	MinimumAmountTax           *bool   `url:"minimumAmountTax,omitempty"`
	Priority                   int     `url:"priority,omitempty"`
	PriorityStart              int     `url:"priorityStart,omitempty"`
	PriorityEnd                int     `url:"priorityEnd,omitempty"`
	ProductRestriction         *bool   `url:"productRestriction,omitempty"`
	Quantity                   int     `url:"quantity,omitempty"`
	QuantityStart              int     `url:"quantityStart,omitempty"`
	QuantityEnd                int     `url:"quantityEnd,omitempty"`
	QuantityPerUser            int     `url:"quantityPerUser,omitempty"`
	QuantityPerUserStart       int     `url:"quantityPerUserStart,omitempty"`
	QuantityPerUserEnd         int     `url:"quantityPerUserEnd,omitempty"`
	ReductionAmount            float64 `url:"reductionAmount,omitempty"`
	ReductionAmountStart       float64 `url:"reductionAmountStart,omitempty"`
	ReductionAmountEnd         float64 `url:"reductionAmountEnd,omitempty"`
	ReductionCurrency          int     `url:"reductionCurrency,omitempty"`
	ReductionCurrencyStart     int     `url:"reductionCurrencyStart,omitempty"`
	ReductionCurrencyEnd       int     `url:"reductionCurrencyEnd,omitempty"`
	ReductionExcludeSpecial    *bool   `url:"reductionExcludeSpecial,omitempty"`
	ReductionPercent           float64 `url:"reductionPercent,omitempty"`
	ReductionPercentStart      float64 `url:"reductionPercentStart,omitempty"`
	ReductionPercentEnd        float64 `url:"reductionPercentEnd,omitempty"`
	ReductionProduct           int     `url:"reductionProduct,omitempty"`
	ReductionProductStart      int     `url:"reductionProductStart,omitempty"`
	ReductionProductEnd        int     `url:"reductionProductEnd,omitempty"`
	ReductionTax               *bool   `url:"reductionTax,omitempty"`
	ShopRestriction            *bool   `url:"shopRestriction,omitempty"`
	To                         string  `url:"to,omitempty"`
	ToStart                    string  `url:"toStart,omitempty"`
	ToEnd                      string  `url:"toEnd,omitempty"`
	GenericListParams
}

// List cart rules. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/CartRules_getPaginatedItems
func (service *CartRulesService) List(shopID string, opts *CartRuleListParams) (*CartRulesResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/cart-rules", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	cartProducts := new(CartRulesResponse)
	response, err := service.client.Do(req, cartProducts)

	if err != nil {
		return nil, response, err
	}

	return cartProducts, response, nil
}

// Get a cart rule. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/CartRules_getSingleItem
func (service *CartRulesService) Get(shopID string, shopContentId string) (*CartRulesResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/cart-rules/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	cartProduct := new(CartRulesResponse)
	response, err := service.client.Do(req, cartProduct)

	if err != nil {
		return nil, response, err
	}

	return cartProduct, response, nil
}
