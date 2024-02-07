package prestashop

import (
	"fmt"
	"net/http"
)

// Category service
type CategoriesService service

type CategoriesResponse struct {
	Category   *Category   `json:"item,omitempty"`
	Categories *[]Category `json:"items,omitempty"`
	GenericListResponse
}

type Category struct {
	ID               string `json:"id,omitempty"`
	ShopID           string `json:"shopId,omitempty"`
	CreatedAt        string `json:"createdAt,omitempty"`
	UpdatedAt        string `json:"updatedAt,omitempty"`
	LastSyncedAt     string `json:"lastSyncedAt,omitempty"`
	IDCategory       int    `json:"idCategory,omitempty"`
	IDParent         int    `json:"idParent,omitempty"`
	UniqueCategoryID string `json:"uniqueCategoryId,omitempty"`
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

type CategoryListParams struct {
	IdCategory             int     `url:"idCategory,omitempty"`
	IdCategoryStart        int     `url:"idCategoryStart,omitempty"`
	IdCategoryEnd          int     `url:"idCategoryEnd,omitempty"`
	IdParent               int     `url:"idParent,omitempty"`
	IdParentStart          int     `url:"idParentStart,omitempty"`
	IdParentEnd            int     `url:"idParentEnd,omitempty"`
	IsoCode                string  `url:"isoCode,omitempty"`
	LinkRewrite            string  `url:"linkRewrite,omitempty"`
	MetaDescription        string  `url:"metaDescription,omitempty"`
	MetaKeywords           string  `url:"metaKeywords,omitempty"`
	MetaTitle              string  `url:"metaTitle,omitempty"`
	Name                   string  `url:"name,omitempty"`
	UniqueCategoryId       string  `url:"uniqueCategoryId,omitempty"`
	Deleted                *bool   `url:"deleted,omitempty"`
	Description            string  `url:"description,omitempty"`
	GenericListParams
}

// List categories. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Categories_getPaginatedItems
func (service *CategoriesService) List(shopID string, opts *CategoryListParams) (*CategoriesResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/categories", service.client.getResourceTypeRaw(), shopID)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	categories := new(CategoriesResponse)
	response, err := service.client.Do(req, categories)

	if err != nil {
		return nil, response, err
	}

	return categories, response, nil
}

// Get a category. Reference: https://docs.cloudsync.prestashop.com/api-doc/expose-raw-api#/operations/Categories_getSingleItem
func (service *CategoriesService) Get(shopID string, shopContentId string) (*CategoriesResponse, *http.Response, error) {
	_url := fmt.Sprintf("%s/%s/categories/%s", service.client.getResourceTypeRaw(), shopID, shopContentId)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	category := new(CategoriesResponse)
	response, err := service.client.Do(req, category)

	if err != nil {
		return nil, response, err
	}

	return category, response, nil
}
