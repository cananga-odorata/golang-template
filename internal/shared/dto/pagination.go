package dto

import "github.com/cananga-odorata/golang-template/internal/shared/domain"

// PaginationRequest represents pagination query parameters
type PaginationRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

// ToPagination converts request to domain pagination
func (r PaginationRequest) ToPagination() domain.Pagination {
	return domain.NewPagination(r.Page, r.PageSize)
}

// PaginatedResponse wraps paginated data
type PaginatedResponse[T any] struct {
	Data       []T   `json:"data"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// NewPaginatedResponse creates a paginated response
func NewPaginatedResponse[T any](data []T, pagination domain.Pagination) PaginatedResponse[T] {
	return PaginatedResponse[T]{
		Data:       data,
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		Total:      pagination.Total,
		TotalPages: pagination.TotalPages(),
	}
}
