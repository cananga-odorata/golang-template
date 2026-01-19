package middleware

import (
	"net/http"
	"strings"

	"github.com/cananga-odorata/golang-template/internal/shared/dto"
	"github.com/cananga-odorata/golang-template/internal/shared/utils"
)

// JWTAuth creates a JWT authentication middleware
func JWTAuth(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				dto.Error(w, http.StatusUnauthorized, "UNAUTHORIZED", "Missing authorization header")
				return
			}

			// Extract Bearer token
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				dto.Error(w, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid authorization header format")
				return
			}

			token := parts[1]

			// TODO: In production, validate JWT properly
			// For now, just check if token starts with "jwt_"
			if !strings.HasPrefix(token, "jwt_") {
				dto.Error(w, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid token")
				return
			}

			// Extract user ID from token (simplified)
			// In production, decode JWT and extract claims
			tokenParts := strings.Split(token, "_")
			if len(tokenParts) < 2 {
				dto.Error(w, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid token format")
				return
			}
			userID := tokenParts[1]

			// Add user info to context
			ctx := utils.SetUserID(r.Context(), userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// OptionalJWTAuth creates a middleware that extracts JWT if present but doesn't require it
func OptionalJWTAuth(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				parts := strings.Split(authHeader, " ")
				if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
					token := parts[1]
					if strings.HasPrefix(token, "jwt_") {
						tokenParts := strings.Split(token, "_")
						if len(tokenParts) >= 2 {
							ctx := utils.SetUserID(r.Context(), tokenParts[1])
							r = r.WithContext(ctx)
						}
					}
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
