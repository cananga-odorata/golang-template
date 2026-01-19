package domain

import "github.com/cananga-odorata/golang-template/internal/shared/domain"

// User domain errors
var (
	ErrUserNotFound = domain.NewNotFoundError("user not found")
	ErrEmailExists  = domain.NewConflictError("email already exists")
	ErrInvalidEmail = domain.NewValidationError("invalid email")
	ErrInvalidRole  = domain.NewValidationError("invalid role")
)
