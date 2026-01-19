package domain

import (
	"time"

	sharedDomain "github.com/cananga-odorata/golang-template/internal/shared/domain"
)

// Affiliate represents an affiliate partner
type Affiliate struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Code      string    `json:"code" db:"code"` // Unique affiliate code
	TenantID  string    `json:"tenant_id" db:"tenant_id"`
	Status    Status    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// AffiliateLink represents a trackable affiliate link
type AffiliateLink struct {
	ID          string    `json:"id" db:"id"`
	AffiliateID string    `json:"affiliate_id" db:"affiliate_id"`
	ProductID   string    `json:"product_id" db:"product_id"`
	URL         string    `json:"url" db:"url"`
	Clicks      int64     `json:"clicks" db:"clicks"`
	Conversions int64     `json:"conversions" db:"conversions"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Status represents affiliate status
type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusPending  Status = "pending"
)

// NewAffiliate creates a new Affiliate
func NewAffiliate(userID, code, tenantID string) *Affiliate {
	now := time.Now()
	return &Affiliate{
		ID:        sharedDomain.NewID(),
		UserID:    userID,
		Code:      code,
		TenantID:  tenantID,
		Status:    StatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
