package domain

// Pagination is the representation of fetch methods return
type Pagination[T any] struct {
	Items T     `json:"items"`
	Total int32 `json:"total"`
}
