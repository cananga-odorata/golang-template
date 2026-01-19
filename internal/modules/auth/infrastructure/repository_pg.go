package infrastructure

import (
	"context"
	"database/sql"

	"github.com/cananga-odorata/golang-template/internal/infra/database"
	"github.com/cananga-odorata/golang-template/internal/modules/auth/domain"
	"github.com/jmoiron/sqlx"
)

type postgresAuthRepository struct {
	db *sqlx.DB
}

// NewPostgresAuthRepository creates a new PostgreSQL auth repository
func NewPostgresAuthRepository(db *sqlx.DB) domain.AuthRepository {
	return &postgresAuthRepository{db: db}
}

func (r *postgresAuthRepository) getQueryable(ctx context.Context) database.Queryable {
	return database.GetQueryable(ctx, r.db)
}

func (r *postgresAuthRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, email, password_hash FROM users WHERE email = $1`
	q := r.getQueryable(ctx)
	err := q.GetContext(ctx, &user, query, email)
	if err == sql.ErrNoRows {
		return nil, domain.ErrUserNotFound
	}
	return &user, err
}

func (r *postgresAuthRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (id, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())`
	q := r.getQueryable(ctx)
	_, err := q.ExecContext(ctx, query, user.ID, user.Email, user.Password)
	return err
}

func (r *postgresAuthRepository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, email, password_hash FROM users WHERE id = $1`
	q := r.getQueryable(ctx)
	err := q.GetContext(ctx, &user, query, id)
	if err == sql.ErrNoRows {
		return nil, domain.ErrUserNotFound
	}
	return &user, err
}
