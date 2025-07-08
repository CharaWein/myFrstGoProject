package repository

import (
	"context"
	"database/sql"
	"fmt"
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
	const query = `
		INSERT INTO profiles (friends, subscribes_count, user_id) 
		VALUES ($1, $2, $3) 
		RETURNING profile_id
	`

	var id uuid.UUID
	if err := r.db.QueryRowContext(ctx, query,
		pq.Array(profile.Friends),
		profile.SubscribesCount,
		profile.UserID,
	).Scan(&id); err != nil {
		return uuid.Nil, fmt.Errorf("failed to create profile: %w", err)
	}

	return id, nil
}

func (r *ProfileRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Profile, error) {
	const query = `
		SELECT profile_id, friends, subscribes_count, user_id 
		FROM profiles 
		WHERE profile_id = $1
	`

	var profile model.Profile
	if err := r.db.QueryRowContext(ctx, query, id).Scan(
		&profile.ProfileID,
		pq.Array(&profile.Friends),
		&profile.SubscribesCount,
		&profile.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("profile not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	return &profile, nil
}

func (r *ProfileRepository) GetAll(ctx context.Context) ([]model.Profile, error) {
	const query = `
		SELECT profile_id, friends, subscribes_count, user_id 
		FROM profiles
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query profiles: %w", err)
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
			return nil, fmt.Errorf("failed to scan profile: %w", err)
		}
		profiles = append(profiles, profile)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return profiles, nil
}

func (r *ProfileRepository) Update(ctx context.Context, id uuid.UUID, profile *model.ProfileUpdate) error {
	const query = `
		UPDATE profiles 
		SET 
			friends = COALESCE($1, friends),
			subscribes_count = COALESCE($2, subscribes_count)
		WHERE profile_id = $3
	`

	var friends interface{}
	if profile.Friends != nil {
		friends = pq.Array(*profile.Friends)
	} else {
		friends = nil
	}

	result, err := r.db.ExecContext(ctx, query,
		friends,
		profile.SubscribesCount,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("profile not found")
	}

	return nil
}

func (r *ProfileRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const query = `
		DELETE FROM profiles 
		WHERE profile_id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("profile not found")
	}

	return nil
}

func (r *ProfileRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*model.Profile, error) {
	const query = `
		SELECT profile_id, friends, subscribes_count, user_id 
		FROM profiles 
		WHERE user_id = $1
	`

	var profile model.Profile
	if err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&profile.ProfileID,
		pq.Array(&profile.Friends),
		&profile.SubscribesCount,
		&profile.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Возвращаем nil вместо ошибки если профиль не найден
		}
		return nil, fmt.Errorf("failed to get profile by user ID: %w", err)
	}

	return &profile, nil
}

func (r *ProfileRepository) SearchByUsername(ctx context.Context, username string) (*model.Profile, error) {
	const query = `
		SELECT p.profile_id, p.friends, p.subscribes_count, p.user_id
		FROM profiles p
		JOIN users u ON p.user_id = u.user_id
		WHERE u.username = $1
	`

	var profile model.Profile
	if err := r.db.QueryRowContext(ctx, query, username).Scan(
		&profile.ProfileID,
		pq.Array(&profile.Friends),
		&profile.SubscribesCount,
		&profile.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to search profile by username: %w", err)
	}

	return &profile, nil
}
