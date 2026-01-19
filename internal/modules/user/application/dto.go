package application

import (
	"time"

	"github.com/cananga-odorata/golang-template/internal/modules/user/domain"
)

// CreateUserRequest DTO for creating a user
type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Role      string `json:"role" validate:"required,oneof=admin affiliate user"`
}

// UpdateUserRequest DTO for updating a user
type UpdateUserRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Status    *string `json:"status,omitempty" validate:"omitempty,oneof=active inactive pending"`
}

// UserResponse DTO for user response
type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	TenantID  string    `json:"tenant_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToUserResponse converts domain User to UserResponse DTO
func ToUserResponse(u *domain.User) *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Role:      string(u.Role),
		Status:    string(u.Status),
		TenantID:  u.TenantID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// ToUserResponseList converts a slice of domain Users to UserResponse DTOs
func ToUserResponseList(users []*domain.User) []*UserResponse {
	responses := make([]*UserResponse, len(users))
	for i, u := range users {
		responses[i] = ToUserResponse(u)
	}
	return responses
}
