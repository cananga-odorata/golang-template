package utils

import "context"

type contextKey string

const (
	userIDKey   contextKey = "user_id"
	tenantIDKey contextKey = "tenant_id"
	userRoleKey contextKey = "user_role"
)

// SetUserID sets the user ID in context
func SetUserID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}

// GetUserID retrieves the user ID from context
func GetUserID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userIDKey).(string)
	return id, ok
}

// MustGetUserID retrieves the user ID from context, panics if not found
func MustGetUserID(ctx context.Context) string {
	id, ok := GetUserID(ctx)
	if !ok {
		panic("user ID not found in context")
	}
	return id
}

// SetTenantID sets the tenant ID in context
func SetTenantID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, tenantIDKey, id)
}

// GetTenantID retrieves the tenant ID from context
func GetTenantID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(tenantIDKey).(string)
	return id, ok
}

// MustGetTenantID retrieves the tenant ID from context, panics if not found
func MustGetTenantID(ctx context.Context) string {
	id, ok := GetTenantID(ctx)
	if !ok {
		panic("tenant ID not found in context")
	}
	return id
}

// SetUserRole sets the user role in context
func SetUserRole(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, userRoleKey, role)
}

// GetUserRole retrieves the user role from context
func GetUserRole(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(userRoleKey).(string)
	return role, ok
}
