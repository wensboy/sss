package model

type Pagination[T any] struct {
	PageBase int   `json:"page_base"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
	Items    []T   `json:"items"`
}
