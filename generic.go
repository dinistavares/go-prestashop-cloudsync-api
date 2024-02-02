package prestashop

type GenericListResponse struct {
	Kind       string                `json:"kind,omitempty"`
	Total      int                   `json:"total,omitempty"`
	Pagination *GenericListPagination `json:"pagination,omitempty"`
	Links      *GenericListLinks      `json:"links,omitempty"`
}

type GenericListPagination struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type GenericListLinks struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
}

type GenericListParams struct {
	Limit  int `url:"limit"`
	Offset int `url:"offset"`
}

var (
	DefaultGenericListParams GenericListParams = GenericListParams{Limit: 10, Offset: 0}
)
