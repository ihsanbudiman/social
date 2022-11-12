package user_usecase

import (
	"context"
	"errors"
	"social/domain"
	"social/helper"
)

type userUseCase struct {
	pgRepo domain.UserRepoPG
}

// Register implements domain.UserUseCase
func (u userUseCase) Register(ctx context.Context, user domain.User) (domain.User, error) {

	// validate email
	isValidEmail := helper.IsValidEmail(user.Email)
	if !isValidEmail {
		return domain.User{}, errors.New("invalid email")
	}

	// hash password
	pwd := helper.GetPwd(user.Password)
	hashedPassword := helper.PashAndSalt(pwd)

	user.Password = hashedPassword

	err := u.pgRepo.RegisterUser(ctx, &user)
	if err != nil {
		return domain.User{}, err
	}

	compare := helper.ComparePasswords(user.Password, pwd)
	if !compare {
		return domain.User{}, errors.New("failed on compare password")
	}

	return user, nil
}

// CheckLogin implements domain.UserUseCase
func (u userUseCase) LoginByEmail(ctx context.Context, email string, password string) (*domain.User, error) {

	isValidEmail := helper.IsValidEmail(email)
	if !isValidEmail {
		return nil, errors.New("invalid email")
	}

	pwd := helper.GetPwd(password)
	// hashedPassword := helper.PashAndSalt(pwd)

	user, err := u.pgRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// compare password
	isValidPassword := helper.ComparePasswords(user.Password, pwd)
	if !isValidPassword {
		return nil, errors.New("username atau password salah")
	}

	return user, nil
}

func (u userUseCase) UpdateUser(ctx context.Context, user domain.User) error {

	// validate email
	isValidEmail := helper.IsValidEmail(user.Email)
	if !isValidEmail {
		return errors.New("invalid email")
	}

	err := u.pgRepo.UpdateUser(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func NewUserUseCase(pgRepo domain.UserRepoPG) domain.UserUseCase {
	return &userUseCase{
		pgRepo: pgRepo,
	}
}
