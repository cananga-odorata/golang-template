package auth

// Domain Entity
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// DTOs
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	User        User   `json:"user"`
}
