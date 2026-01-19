package domain

import (
	"time"

	"github.com/google/uuid"
)

// BaseEntity provides common fields for all entities
type BaseEntity struct {
	ID        string    `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// NewID generates a new UUID
func NewID() string {
	return uuid.New().String()
}

// NewBaseEntity creates a new BaseEntity with generated ID and timestamps
func NewBaseEntity() BaseEntity {
	now := time.Now()
	return BaseEntity{
		ID:        NewID(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}
