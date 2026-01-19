package application

import (
	"context"
	"time"

	"github.com/cananga-odorata/golang-template/internal/modules/user/domain"
	"github.com/cananga-odorata/golang-template/internal/shared/utils"
	"golang.org/x/crypto/bcrypt"
)

// UserService defines the user business logic interface
type UserService interface {
	Create(ctx context.Context, req CreateUserRequest) (*UserResponse, error)
	GetByID(ctx context.Context, id string) (*UserResponse, error)
	GetByEmail(ctx context.Context, email string) (*UserResponse, error)
	Update(ctx context.Context, id string, req UpdateUserRequest) (*UserResponse, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter domain.UserFilter) ([]*UserResponse, int64, error)
}

type userService struct {
	repo domain.UserRepository
}

// NewUserService creates a new UserService with dependencies
func NewUserService(repo domain.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(ctx context.Context, req CreateUserRequest) (*UserResponse, error) {
	// Check if email exists
	existing, _ := s.repo.GetByEmail(ctx, req.Email)
	if existing != nil {
		return nil, domain.ErrEmailExists
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Extract tenant from context (for multi-tenancy)
	tenantID, _ := utils.GetTenantID(ctx)

	// Create domain entity
	user, err := domain.NewUser(
		req.Email,
		string(hash),
		req.FirstName,
		req.LastName,
		domain.Role(req.Role),
		tenantID,
	)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return ToUserResponse(user), nil
}

func (s *userService) GetByID(ctx context.Context, id string) (*UserResponse, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToUserResponse(user), nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*UserResponse, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return ToUserResponse(user), nil
}

func (s *userService) Update(ctx context.Context, id string, req UpdateUserRequest) (*UserResponse, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Apply updates
	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}
	if req.Status != nil {
		user.Status = domain.Status(*req.Status)
	}
	user.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return ToUserResponse(user), nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	// Verify user exists
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}

func (s *userService) List(ctx context.Context, filter domain.UserFilter) ([]*UserResponse, int64, error) {
	// Auto-filter by tenant if available
	if tenantID, ok := utils.GetTenantID(ctx); ok && filter.TenantID == "" {
		filter.TenantID = tenantID
	}

	users, total, err := s.repo.List(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return ToUserResponseList(users), total, nil
}
