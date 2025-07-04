package model

import "github.com/google/uuid"

type Profile struct {
    ProfileID       uuid.UUID `json:"profile_id"`
    Friends        []string  `json:"friends"`
    SubscribesCount int       `json:"subscribes_count"`
    UserID         uuid.UUID `json:"user_id"`
}

type ProfileCreate struct {
    Friends        []string  `json:"friends"`
    SubscribesCount int       `json:"subscribes_count"`
    UserID         uuid.UUID `json:"user_id" validate:"required"`
}

type ProfileUpdate struct {
    Friends        *[]string `json:"friends,omitempty"`
    SubscribesCount *int      `json:"subscribes_count,omitempty"`
}