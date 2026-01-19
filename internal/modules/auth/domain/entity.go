package domain

import (
	"time"

	sharedDomain "github.com/cananga-odorata/golang-template/internal/shared/domain"
)

// User represents the authenticated user entity
type User struct {
	ID       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"-" db:"password_hash"`
}

// Session represents a user session
type Session struct {
	ID           string    `json:"id" db:"id"`
	UserID       string    `json:"user_id" db:"user_id"`
	RefreshToken string    `json:"-" db:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// NewSession creates a new session
func NewSession(userID, refreshToken string, expiresAt time.Time) *Session {
	return &Session{
		ID:           sharedDomain.NewID(),
		UserID:       userID,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
		CreatedAt:    time.Now(),
	}
}
