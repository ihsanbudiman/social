package domain

import (
	"context"
	"encoding/json"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// remove password when marshaling a json
func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Password string `json:"password,omitempty"`
		*Alias
	}{
		Password: "",
		Alias:    (*Alias)(u),
	})
}

type UserRepo interface {
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByEmailAndUsername(ctx context.Context, email, username string) (User, error)
	GetByID(ctx context.Context, id uint) (User, error)
	UpdateUser(ctx context.Context, user *User) error
	RegisterUser(ctx context.Context, user *User) error
}

type UserUseCase interface {
	LoginByEmail(ctx context.Context, email, password string) (User, error)
	Register(ctx context.Context, user User) (User, error)
}
