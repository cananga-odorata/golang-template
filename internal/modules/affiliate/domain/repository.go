package domain

import "context"

// AffiliateRepository defines the contract for affiliate persistence
type AffiliateRepository interface {
	Create(ctx context.Context, affiliate *Affiliate) error
	GetByID(ctx context.Context, id string) (*Affiliate, error)
	GetByUserID(ctx context.Context, userID string) (*Affiliate, error)
	GetByCode(ctx context.Context, code string) (*Affiliate, error)
	Update(ctx context.Context, affiliate *Affiliate) error
	List(ctx context.Context, filter AffiliateFilter) ([]*Affiliate, int64, error)
}

// AffiliateLinkRepository defines the contract for affiliate link persistence
type AffiliateLinkRepository interface {
	Create(ctx context.Context, link *AffiliateLink) error
	GetByID(ctx context.Context, id string) (*AffiliateLink, error)
	IncrementClicks(ctx context.Context, id string) error
	IncrementConversions(ctx context.Context, id string) error
	ListByAffiliateID(ctx context.Context, affiliateID string) ([]*AffiliateLink, error)
}

// AffiliateFilter holds filters for listing affiliates
type AffiliateFilter struct {
	TenantID string
	Status   *Status
	Limit    int
	Offset   int
}
