package model

import "github.com/google/uuid"

type User struct {
    UserID    uuid.UUID `json:"user_id"`
    Email     string    `json:"email" validate:"required,email"`
    Username  string    `json:"username" validate:"required"`
    Messages  []string  `json:"messages"`
}

type UserCreate struct {
    Email    string `json:"email" validate:"required,email"`
    Username string `json:"username" validate:"required"`
}

type UserUpdate struct {
    Email    *string `json:"email,omitempty" validate:"omitempty,email"`
    Username *string `json:"username,omitempty" validate:"omitempty"`
    Messages []string `json:"messages,omitempty"`
}