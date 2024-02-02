package prestashop

type GenericListResponse struct {
	Kind       string                 `json:"kind,omitempty"`
	Total      int                    `json:"total,omitempty"`
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
	Limit             int    `url:"limit"`
	Offset            int    `url:"offset"`
	CreatedAt         string `url:"createdAt,omitempty"`
	CreatedAtStart    string `url:"createdAtStart,omitempty"`
	CreatedAtEnd      string `url:"createdAtEnd,omitempty"`
	Direction         string `url:"direction,omitempty"`
	LastSyncedAt      string `url:"lastSyncedAt,omitempty"`
	LastSyncedAtStart string `url:"lastSyncedAtStart,omitempty"`
	LastSyncedAtEnd   string `url:"lastSyncedAtEnd,omitempty"`
	Sort              string `url:"sort,omitempty"`
	UpdatedAt         string `url:"updatedAt,omitempty"`
	UpdatedAtStart    string `url:"updatedAtStart,omitempty"`
	UpdatedAtEnd      string `url:"updatedAtEnd,omitempty"`
}

var (
	DefaultGenericListParams GenericListParams = GenericListParams{Limit: 10, Offset: 0}
)
