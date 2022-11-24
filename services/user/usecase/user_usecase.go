package user_usecase

import (
	"context"
	"errors"
	"social/domain"
	"social/helper"
	"social/opentelemetry"

	"gorm.io/gorm"
)

type userUseCase struct {
	pgRepo domain.UserRepo
}

// Register implements domain.UserUseCase
func (u userUseCase) Register(ctx context.Context, user domain.User) (domain.User, error) {

	if user.Email == "" {
		return domain.User{}, errors.New("email is required")
	}

	// validate email
	isValidEmail := helper.IsValidEmail(user.Email)
	if !isValidEmail {
		return domain.User{}, errors.New("invalid email")
	}

	if user.Name == "" {
		return domain.User{}, errors.New("name is required")
	}

	if user.Password == "" {
		return domain.User{}, errors.New("password is required")
	}

	// check password is strong enough
	if isStrongPassword := helper.IsStrongPassword(user.Password); !isStrongPassword {
		return domain.User{}, errors.New("password is not strong enough, minimum 8 characters, at least one uppercase letter, one lowercase letter, one number and one special character")
	}

	if user.Username == "" {
		return domain.User{}, errors.New("username is required")
	}

	// check if email already exists
	otherUser, err := u.pgRepo.GetByEmailAndUsername(ctx, user.Email, user.Username)

	if err != nil && err != gorm.ErrRecordNotFound {
		return domain.User{}, err
	}

	if otherUser.ID != 0 {
		return domain.User{}, errors.New("email or username already exists")
	}

	// hash password
	pwd := helper.GetPwd(user.Password)
	hashedPassword := helper.HashAndSalt(pwd)

	user.Password = hashedPassword

	err = u.pgRepo.RegisterUser(ctx, &user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// CheckLogin implements domain.UserUseCase
func (u userUseCase) LoginByEmail(ctx context.Context, email string, password string) (domain.User, error) {
	tracer := opentelemetry.GetTracer()

	ucaseCtx, span := tracer.Start(ctx, "user_usecase.LoginByEmail")
	defer span.End()

	isValidEmail := helper.IsValidEmail(email)
	if !isValidEmail {
		return domain.User{}, errors.New("invalid email")
	}

	_, spanGetPWD := tracer.Start(ucaseCtx, "user_usecase.LoginByEmail.GetPWD")
	pwd := helper.GetPwd(password)
	spanGetPWD.End()

	// hashedPassword := helper.HashAndSalt(pwd)

	user, err := u.pgRepo.GetByEmail(ucaseCtx, email)
	if err != nil {
		return domain.User{}, err
	}

	// compare password
	_, spanComparePassword := tracer.Start(ucaseCtx, "user_usecase.LoginByEmail.ComparePassword")
	isValidPassword := helper.ComparePasswords(user.Password, pwd)
	spanComparePassword.End()

	if !isValidPassword {
		return domain.User{}, errors.New("username atau password salah")
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

func NewUserUseCase(pgRepo domain.UserRepo) domain.UserUseCase {
	return &userUseCase{
		pgRepo: pgRepo,
	}
}
