package prestashop

import (
	"fmt"
	"net/http"
	"time"
)

// Customer service
type CustomersService service

type CustomersResponse struct {
	Customer  Customer   `json:"item,omitempty"`
	Customers []Customer `json:"items,omitempty"`
	GenericListResponse
}

type Customer struct {
	IDCustomer        int        `json:"idCustomer,omitempty"`
	IDLang            int        `json:"idLang,omitempty"`
	ID                string     `json:"id,omitempty"`
	ShopID            string     `json:"shopId,omitempty"`
	EmailHash         string     `json:"emailHash,omitempty"`
	Newsletter        *bool      `json:"newsletter,omitempty"`
	Optin             *bool      `json:"optin,omitempty"`
	Active            *bool      `json:"active,omitempty"`
	IsGuest           *bool      `json:"isGuest,omitempty"`
	Deleted           *bool      `json:"deleted,omitempty"`
	CreatedAt         *time.Time `json:"createdAt,omitempty"`
	UpdatedAt         *time.Time `json:"updatedAt,omitempty"`
	LastSyncedAt      *time.Time `json:"lastSyncedAt,omitempty"`
	NewsletterDateAdd *time.Time `json:"newsletterDateAdd,omitempty"`
	DeletedAt         *time.Time `json:"deletedAt,omitempty"`
}

type CustomerListParams struct {
	DeletedAt              string `url:"deletedAt,omitempty"`
	DeletedAtStart         string `url:"deletedAtStart,omitempty"`
	DeletedAtEnd           string `url:"deletedAtEnd,omitempty"`
	EmailHash              string `url:"emailHash,omitempty"`
	IdCustomer             *int   `url:"idCustomer,omitempty"`
	IdCustomerStart        *int   `url:"idCustomerStart,omitempty"`
	IdCustomerEnd          *int   `url:"idCustomerEnd,omitempty"`
	IdLang                 *int   `url:"idLang,omitempty"`
	IdLangStart            *int   `url:"idLangStart,omitempty"`
	IdLangEnd              *int   `url:"idLangEnd,omitempty"`
	Active                 *bool  `url:"active,omitempty"`
	Deleted                *bool  `url:"deleted,omitempty"`
	IsGuest                *bool  `url:"isGuest,omitempty"`
	Optin                  *bool  `url:"optin,omitempty"`
	Newsletter             *bool  `url:"newsletter,omitempty"`
	NewsletterDateAdd      string `url:"newsletterDateAdd,omitempty"`
	NewsletterDateAddStart string `url:"newsletterDateAddStart,omitempty"`
	NewsletterDateAddEnd   string `url:"newsletterDateAddEnd,omitempty"`
	GenericListParams
}

// List customers. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Customers_getPaginatedItems
func (service *CustomersService) List(shopID string, opts *CustomerListParams) (*CustomersResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/customers", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	customers := new(CustomersResponse)
	response, err := service.client.Do(req, customers)

	if err != nil {
		return nil, response, err
	}

	return customers, response, nil
}

// Get a customer. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Customers_getSingleItem
func (service *CustomersService) Get(shopID string, shopContentId string) (*CustomersResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/customers/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	customer := new(CustomersResponse)
	response, err := service.client.Do(req, customer)

	if err != nil {
		return nil, response, err
	}

	return customer, response, nil
}
