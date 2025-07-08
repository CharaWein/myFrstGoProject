package model

import "github.com/google/uuid"

type Profile struct {
	ProfileID       uuid.UUID `json:"profile_id"`
	Friends         []string  `json:"friends"`
	SubscribesCount int       `json:"subscribes_count"`
	UserID          uuid.UUID `json:"user_id"`
}

type ProfileCreate struct {
	Friends         []string  `json:"friends" validate:"required"`
	SubscribesCount int       `json:"subscribes_count" validate:"required,min=0"`
	UserID          uuid.UUID `json:"user_id" validate:"required"`
}

type ProfileUpdate struct {
	Friends         *[]string `json:"friends,omitempty" validate:"omitempty"`
	SubscribesCount *int      `json:"subscribes_count,omitempty" validate:"omitempty,min=0"`
}
