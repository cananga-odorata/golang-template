package domain

// Pagination holds pagination parameters
type Pagination struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

// NewPagination creates a new Pagination with defaults
func NewPagination(page, pageSize int) Pagination {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}
	return Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}

// Offset returns the offset for database queries
func (p Pagination) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// Limit returns the limit for database queries
func (p Pagination) Limit() int {
	return p.PageSize
}

// TotalPages returns the total number of pages
func (p Pagination) TotalPages() int {
	if p.Total == 0 {
		return 0
	}
	pages := int(p.Total) / p.PageSize
	if int(p.Total)%p.PageSize > 0 {
		pages++
	}
	return pages
}

// HasNext returns true if there is a next page
func (p Pagination) HasNext() bool {
	return p.Page < p.TotalPages()
}

// HasPrev returns true if there is a previous page
func (p Pagination) HasPrev() bool {
	return p.Page > 1
}
