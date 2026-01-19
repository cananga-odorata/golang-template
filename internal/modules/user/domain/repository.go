package domain

import "context"

// UserRepository defines the contract for user persistence
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter UserFilter) ([]*User, int64, error)
}

// UserFilter holds filters for listing users
type UserFilter struct {
	TenantID string
	Role     *Role
	Status   *Status
	Search   string
	Limit    int
	Offset   int
}
