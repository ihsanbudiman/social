package domain

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex" json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserRepo interface {
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id uint) (User, error)
	UpdateUser(ctx context.Context, user *User) error
	RegisterUser(ctx context.Context, user *User) error
}

type UserUseCase interface {
	LoginByEmail(ctx context.Context, email, password string) (*User, error)
	Register(ctx context.Context, user User) (User, error)
}
