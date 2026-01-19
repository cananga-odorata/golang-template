package domain

import "context"

// AuthRepository defines the contract for auth persistence
type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id string) (*User, error)
}

// SessionRepository defines the contract for session persistence
type SessionRepository interface {
	Create(ctx context.Context, session *Session) error
	GetByRefreshToken(ctx context.Context, token string) (*Session, error)
	Delete(ctx context.Context, id string) error
	DeleteByUserID(ctx context.Context, userID string) error
}
