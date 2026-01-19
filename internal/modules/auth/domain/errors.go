package domain

import "github.com/cananga-odorata/golang-template/internal/shared/domain"

// Auth domain errors
var (
	ErrUserNotFound       = domain.NewNotFoundError("user not found")
	ErrInvalidCredentials = domain.NewUnauthorizedError("invalid credentials")
	ErrEmailExists        = domain.NewConflictError("email already exists")
	ErrTokenExpired       = domain.NewUnauthorizedError("token expired")
	ErrInvalidToken       = domain.NewUnauthorizedError("invalid token")
)
