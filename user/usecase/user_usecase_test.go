package user_usecase

import (
	"context"
	"errors"
	"fmt"
	"social/domain"
	"social/helper"
	user_repo_pg_mock "social/user/repo/pg/mock"
	"testing"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func Test_userUseCase_CheckLogin(t *testing.T) {

	t.Run("fail email", func(t *testing.T) {
		ctx := context.Background()
		userRepoPGMock := user_repo_pg_mock.NewUserRepoPGMock()

		userUseCase := NewUserUseCase(userRepoPGMock)

		_, err := userUseCase.LoginByEmail(ctx, "email", "password")
		if err == nil {
			t.Errorf("expected error, got nil")
		}

	})

	t.Run("should error on GetByEmail", func(t *testing.T) {
		ctx := context.Background()

		password := "ihsan"

		userRepoPGMock := user_repo_pg_mock.NewUserRepoPGMock()
		userRepoPGMock.On("GetByEmail", ctx, "ihsan@gmail.com").Return(&domain.User{}, errors.New("error"))

		userUseCase := NewUserUseCase(userRepoPGMock)

		_, err := userUseCase.LoginByEmail(ctx, "ihsan@gmail.com", password)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

	})

	t.Run("failed on compare password", func(t *testing.T) {
		ctx := context.Background()

		password := "ihsan"
		pwd := helper.GetPwd(password)
		hashedPassword := helper.PashAndSalt(pwd)

		userRepoPGMock := user_repo_pg_mock.NewUserRepoPGMock()
		userRepoPGMock.On("GetByEmail", ctx, "ihsan@gmail.com").Return(&domain.User{
			Model: gorm.Model{
				ID: 1,
			},
			Name:     "ihsan",
			Email:    "ihsan@gmail.com",
			Password: hashedPassword,
		}, nil)

		userUseCase := NewUserUseCase(userRepoPGMock)

		_, err := userUseCase.LoginByEmail(ctx, "ihsan@gmail.com", "anotherPassword")
		if err == nil {
			t.Errorf("expected error, got nil")
		}

	})

	t.Run("test success", func(t *testing.T) {
		ctx := context.Background()

		password := "ihsan"
		pwd := helper.GetPwd(password)
		hashedPassword := helper.PashAndSalt(pwd)

		userRepoPGMock := user_repo_pg_mock.NewUserRepoPGMock()
		userRepoPGMock.On("GetByEmail", ctx, "ihsan@gmail.com").Return(&domain.User{
			Model: gorm.Model{
				ID: 1,
			},
			Name:     "ihsan",
			Email:    "ihsan@gmail.com",
			Password: hashedPassword,
		}, nil)

		userUseCase := NewUserUseCase(userRepoPGMock)

		_, err := userUseCase.LoginByEmail(ctx, "ihsan@gmail.com", password)
		if err != nil {
			t.Error("error should be nil")
		}
	})
}

func Test_userUseCase_Register(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		userRepoPGMock := user_repo_pg_mock.NewUserRepoPGMock()
		pwd := helper.GetPwd("ihsan123")
		hashedPassword := helper.PashAndSalt(pwd)

		userRepoPGMock.On("RegisterUser", ctx, &domain.User{
			Email:    "ihsan@gmail.com",
			Password: hashedPassword,
			Name:     "ihsan",
		}).Run(func(args mock.Arguments) {
			fmt.Println(234567)
		}).Return(nil)

		userUseCase := NewUserUseCase(userRepoPGMock)

		_, err := userUseCase.Register(ctx, domain.User{
			Email:    "ihsan@gmail.com",
			Password: "ihsan123",
			Name:     "ihsan",
		})

		if err != nil {
			t.Errorf("err %v", err)
		}

	})

}
