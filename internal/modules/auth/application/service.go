package application

import (
	"context"
	"time"

	"github.com/cananga-odorata/golang-template/internal/modules/auth/domain"
	sharedDomain "github.com/cananga-odorata/golang-template/internal/shared/domain"
	"golang.org/x/crypto/bcrypt"
)

// AuthService defines the auth business logic interface
type AuthService interface {
	Register(ctx context.Context, req RegisterRequest) (*AuthResponse, error)
	Login(ctx context.Context, req LoginRequest) (*AuthResponse, error)
	RefreshToken(ctx context.Context, req RefreshTokenRequest) (*AuthResponse, error)
}

type authService struct {
	repo      domain.AuthRepository
	jwtSecret string
}

// NewAuthService creates a new AuthService
func NewAuthService(repo domain.AuthRepository, jwtSecret string) AuthService {
	return &authService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *authService) Register(ctx context.Context, req RegisterRequest) (*AuthResponse, error) {
	// Check if email exists
	existing, _ := s.repo.GetUserByEmail(ctx, req.Email)
	if existing != nil {
		return nil, domain.ErrEmailExists
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &domain.User{
		ID:       sharedDomain.NewID(),
		Email:    req.Email,
		Password: string(hash),
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	// Generate tokens (simplified - in production use JWT)
	accessToken := "jwt_" + user.ID + "_" + time.Now().Format(time.RFC3339)

	return &AuthResponse{
		AccessToken: accessToken,
		ExpiresIn:   3600,
		User: UserResponse{
			ID:    user.ID,
			Email: user.Email,
		},
	}, nil
}

func (s *authService) Login(ctx context.Context, req LoginRequest) (*AuthResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	// Generate tokens (simplified - in production use JWT)
	accessToken := "jwt_" + user.ID + "_" + time.Now().Format(time.RFC3339)
	refreshToken := "refresh_" + user.ID + "_" + time.Now().Format(time.RFC3339)

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    3600,
		User: UserResponse{
			ID:    user.ID,
			Email: user.Email,
		},
	}, nil
}

func (s *authService) RefreshToken(ctx context.Context, req RefreshTokenRequest) (*AuthResponse, error) {
	// In production, validate refresh token from database
	// For now, just return a new access token
	return &AuthResponse{
		AccessToken: "jwt_refreshed_" + time.Now().Format(time.RFC3339),
		ExpiresIn:   3600,
	}, nil
}
