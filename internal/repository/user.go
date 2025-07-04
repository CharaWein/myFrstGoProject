package repository

import (
	"context"
	"database/sql"
	"go-backend-template/internal/model"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.UserCreate) (uuid.UUID, error) {
	var id uuid.UUID
	query := `INSERT INTO users (email, username) VALUES ($1, $2) RETURNING user_id`
	err := r.db.QueryRowContext(ctx, query, user.Email, user.Username).Scan(&id)
	return id, err
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var user model.User
	query := `SELECT user_id, email, username, messages FROM users WHERE user_id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.UserID,
		&user.Email,
		&user.Username,
		pq.Array(&user.Messages),
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	query := `SELECT user_id, email, username, messages FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.UserID,
			&user.Email,
			&user.Username,
			pq.Array(&user.Messages),
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, id uuid.UUID, user *model.UserUpdate) error {
	query := `UPDATE users SET 
              email = COALESCE($1, email),
              username = COALESCE($2, username),
              messages = COALESCE($3, messages)
              WHERE user_id = $4`

	_, err := r.db.ExecContext(ctx, query,
		user.Email,
		user.Username,
		pq.Array(user.Messages),
		id,
	)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	query := `SELECT user_id, email, username, messages FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.UserID,
		&user.Email,
		&user.Username,
		pq.Array(&user.Messages),
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	query := `SELECT user_id, email, username, messages FROM users WHERE username = $1`
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.UserID,
		&user.Email,
		&user.Username,
		pq.Array(&user.Messages),
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
