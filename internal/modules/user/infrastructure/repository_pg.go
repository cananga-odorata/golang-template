package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cananga-odorata/golang-template/internal/infra/database"
	"github.com/cananga-odorata/golang-template/internal/modules/user/domain"
	"github.com/jmoiron/sqlx"
)

type postgresUserRepository struct {
	db *sqlx.DB
}

// NewPostgresUserRepository creates a new PostgreSQL user repository
func NewPostgresUserRepository(db *sqlx.DB) domain.UserRepository {
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) getQueryable(ctx context.Context) database.Queryable {
	return database.GetQueryable(ctx, r.db)
}

func (r *postgresUserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (id, email, password_hash, first_name, last_name, role, status, tenant_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	q := r.getQueryable(ctx)
	_, err := q.ExecContext(ctx, query,
		user.ID, user.Email, user.PasswordHash, user.FirstName, user.LastName,
		user.Role, user.Status, user.TenantID, user.CreatedAt, user.UpdatedAt,
	)
	return err
}

func (r *postgresUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	query := `SELECT * FROM users WHERE id = $1`
	q := r.getQueryable(ctx)
	err := q.GetContext(ctx, &user, query, id)
	if err == sql.ErrNoRows {
		return nil, domain.ErrUserNotFound
	}
	return &user, err
}

func (r *postgresUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT * FROM users WHERE email = $1`
	q := r.getQueryable(ctx)
	err := q.GetContext(ctx, &user, query, email)
	if err == sql.ErrNoRows {
		return nil, domain.ErrUserNotFound
	}
	return &user, err
}

func (r *postgresUserRepository) Update(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE users
		SET first_name = $2, last_name = $3, status = $4, updated_at = $5
		WHERE id = $1
	`
	q := r.getQueryable(ctx)
	_, err := q.ExecContext(ctx, query,
		user.ID, user.FirstName, user.LastName, user.Status, user.UpdatedAt,
	)
	return err
}

func (r *postgresUserRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	q := r.getQueryable(ctx)
	_, err := q.ExecContext(ctx, query, id)
	return err
}

func (r *postgresUserRepository) List(ctx context.Context, filter domain.UserFilter) ([]*domain.User, int64, error) {
	var users []*domain.User
	var total int64

	// Build dynamic query
	whereClause := "WHERE 1=1"
	args := []interface{}{}
	argIdx := 1

	if filter.TenantID != "" {
		whereClause += fmt.Sprintf(" AND tenant_id = $%d", argIdx)
		args = append(args, filter.TenantID)
		argIdx++
	}

	if filter.Role != nil {
		whereClause += fmt.Sprintf(" AND role = $%d", argIdx)
		args = append(args, *filter.Role)
		argIdx++
	}

	if filter.Status != nil {
		whereClause += fmt.Sprintf(" AND status = $%d", argIdx)
		args = append(args, *filter.Status)
		argIdx++
	}

	if filter.Search != "" {
		whereClause += fmt.Sprintf(" AND (first_name ILIKE $%d OR last_name ILIKE $%d OR email ILIKE $%d)", argIdx, argIdx, argIdx)
		args = append(args, "%"+filter.Search+"%")
		argIdx++
	}

	// Count total
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM users %s", whereClause)
	q := r.getQueryable(ctx)
	if err := q.GetContext(ctx, &total, countQuery, args...); err != nil {
		return nil, 0, err
	}

	// Get paginated results
	query := fmt.Sprintf("SELECT * FROM users %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereClause, argIdx, argIdx+1)
	args = append(args, filter.Limit, filter.Offset)

	if err := q.SelectContext(ctx, &users, query, args...); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
