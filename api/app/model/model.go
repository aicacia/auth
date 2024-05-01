package model

type PaginationST[T any] struct {
	HasMore bool `json:"has_more" validate:"required"`
	Items   []T  `json:"items" validate:"required"`
} // @name Pagination
