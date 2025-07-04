package repository

import (
	"context"
	"database/sql"
	"go-backend-template/internal/model"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) Create(ctx context.Context, profile *model.ProfileCreate) (uuid.UUID, error) {
	var id uuid.UUID
	query := `INSERT INTO profiles (friends, subscribes_count, user_id) 
              VALUES ($1, $2, $3) RETURNING profile_id`
	err := r.db.QueryRowContext(ctx, query,
		pq.Array(profile.Friends),
		profile.SubscribesCount,
		profile.UserID).Scan(&id)
	return id, err
}

func (r *ProfileRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Profile, error) {
	var profile model.Profile
	query := `SELECT profile_id, friends, subscribes_count, user_id 
              FROM profiles WHERE profile_id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&profile.ProfileID,
		pq.Array(&profile.Friends),
		&profile.SubscribesCount,
		&profile.UserID,
	)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepository) GetAll(ctx context.Context) ([]model.Profile, error) {
	query := `SELECT profile_id, friends, subscribes_count, user_id FROM profiles`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []model.Profile
	for rows.Next() {
		var profile model.Profile
		if err := rows.Scan(
			&profile.ProfileID,
			pq.Array(&profile.Friends),
			&profile.SubscribesCount,
			&profile.UserID,
		); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

func (r *ProfileRepository) Update(ctx context.Context, id uuid.UUID, profile *model.ProfileUpdate) error {
	query := `UPDATE profiles SET 
              friends = COALESCE($1, friends),
              subscribes_count = COALESCE($2, subscribes_count)
              WHERE profile_id = $3`

	_, err := r.db.ExecContext(ctx, query,
		pq.Array(profile.Friends),
		profile.SubscribesCount,
		id,
	)
	return err
}

func (r *ProfileRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM profiles WHERE profile_id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *ProfileRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*model.Profile, error) {
	var profile model.Profile
	query := `SELECT profile_id, friends, subscribes_count, user_id 
              FROM profiles WHERE user_id = $1`
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&profile.ProfileID,
		pq.Array(&profile.Friends),
		&profile.SubscribesCount,
		&profile.UserID,
	)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
