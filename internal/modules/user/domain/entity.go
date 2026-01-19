package domain

import (
	"time"

	"github.com/cananga-odorata/golang-template/internal/shared/domain"
)

// Role represents user roles
type Role string

const (
	RoleAdmin     Role = "admin"
	RoleAffiliate Role = "affiliate"
	RoleUser      Role = "user"
)

// Status represents user status
type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusPending  Status = "pending"
)

// User represents the user domain entity
type User struct {
	ID           string    `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
	Role         Role      `json:"role" db:"role"`
	Status       Status    `json:"status" db:"status"`
	TenantID     string    `json:"tenant_id" db:"tenant_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// NewUser creates a new User with validation
func NewUser(email, passwordHash, firstName, lastName string, role Role, tenantID string) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if passwordHash == "" {
		return nil, domain.NewValidationError("password is required")
	}

	now := time.Now()
	return &User{
		ID:           domain.NewID(),
		Email:        email,
		PasswordHash: passwordHash,
		FirstName:    firstName,
		LastName:     lastName,
		Role:         role,
		Status:       StatusPending,
		TenantID:     tenantID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}

// FullName returns the user's full name
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// IsActive returns true if user is active
func (u *User) IsActive() bool {
	return u.Status == StatusActive
}

// Activate activates the user
func (u *User) Activate() {
	u.Status = StatusActive
	u.UpdatedAt = time.Now()
}

// Deactivate deactivates the user
func (u *User) Deactivate() {
	u.Status = StatusInactive
	u.UpdatedAt = time.Now()
}
