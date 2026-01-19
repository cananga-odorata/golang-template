package domain

import "context"

// ProductRepository defines the contract for product persistence
type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter ProductFilter) ([]*Product, int64, error)
}

// ProductFilter holds filters for listing products
type ProductFilter struct {
	TenantID string
	Status   *Status
	Search   string
	Limit    int
	Offset   int
}
