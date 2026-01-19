package application

// RegisterRequest DTO for user registration
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// LoginRequest DTO for user login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RefreshTokenRequest DTO for token refresh
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// AuthResponse DTO for auth responses
type AuthResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token,omitempty"`
	ExpiresIn    int          `json:"expires_in"`
	User         UserResponse `json:"user"`
}

// UserResponse DTO for user data in auth responses
type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
