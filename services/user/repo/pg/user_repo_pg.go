package user_repo_pg

import (
	"context"

	"gorm.io/gorm"

	"social/domain"
)

type userRepoPg struct {
	// gorm connection
	db *gorm.DB
}

// RegisterUser implements domain.UserRepoPG
func (u userRepoPg) RegisterUser(ctx context.Context, user *domain.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

// GetByID implements domain.UserRepoPG
func (u userRepoPg) GetByID(ctx context.Context, id uint) (domain.User, error) {
	var user domain.User

	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// UpdateUser implements domain.UserRepoPG
func (u userRepoPg) UpdateUser(ctx context.Context, user *domain.User) error {
	err := u.db.Updates(domain.User{
		Email: user.Email,
		Name:  user.Name,
	}).Where("users.id = ?", user.ID).Error

	if err != nil {
		return err
	}

	return nil
}

// CheckLogin implements domain.UserRepoPG
func (u userRepoPg) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserRepoPG(db *gorm.DB) domain.UserRepo {
	return &userRepoPg{
		db: db,
	}
}
