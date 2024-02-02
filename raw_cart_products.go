package prestashop

import (
	"fmt"
	"net/http"
	"time"
)

// Cart Products service
type CartProductsService service

type CartProductsResponse struct {
	CartProduct  *CartProduct   `json:"item,omitempty"`
	CartProducts *[]CartProduct `json:"items,omitempty"`
	GenericListResponse
}

type CartProduct struct {
	Quantity           int        `json:"quantity,omitempty"`
	ID                 string     `json:"id,omitempty"`
	ShopID             string     `json:"shopId,omitempty"`
	IDCartProduct      string     `json:"idCartProduct,omitempty"`
	IDCart             string     `json:"idCart,omitempty"`
	IDProduct          string     `json:"idProduct,omitempty"`
	IDProductAttribute string     `json:"idProductAttribute,omitempty"`
	CreatedAt          *time.Time `json:"createdAt,omitempty"`
	UpdatedAt          *time.Time `json:"updatedAt,omitempty"`
	LastSyncedAt       *time.Time `json:"lastSyncedAt,omitempty"`
}

type CartProductListParams struct {
	IdCart            string `url:"idCart,omitempty"`
	IdCartProduct     string `url:"idCartProduct,omitempty"`
	IdProduct         string `url:"idProduct,omitempty"`
	ProductAttribute string `url:"idProductAttribute,omitempty"`
	Quantity          int    `url:"quantity,omitempty"`

	GenericListParams
}

// List cartProducts. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/CartProducts_getPaginatedItems
func (service *CartProductsService) List(shopID string, opts *CartProductListParams) (*CartProductsResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/cart-products", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	cartProducts := new(CartProductsResponse)
	response, err := service.client.Do(req, cartProducts)

	if err != nil {
		return nil, response, err
	}

	return cartProducts, response, nil
}

// Get a cartProduct. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/CartProducts_getSingleItem
func (service *CartProductsService) Get(shopID string, shopContentId string) (*CartProductsResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/cart-products/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	cartProduct := new(CartProductsResponse)
	response, err := service.client.Do(req, cartProduct)

	if err != nil {
		return nil, response, err
	}

	return cartProduct, response, nil
}
