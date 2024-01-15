package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;type:uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreate struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUser(user UserCreate) User {
	now := time.Now().UTC()

	return User{
		ID:        uuid.New(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
