package user_usecase

import (
	"context"
	"errors"
	"social/domain"
	"social/helper"
	user_repo_pg_mock "social/services/user/repo/pg/mock"
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
		userRepoPGMock.On("GetByEmail", ctx, "ihsan@gmail.com").Return(domain.User{}, errors.New("error"))

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
		hashedPassword := helper.HashAndSalt(pwd)

		userRepoPGMock := user_repo_pg_mock.NewUserRepoPGMock()
		userRepoPGMock.On("GetByEmail", ctx, "ihsan@gmail.com").Return(domain.User{
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
		hashedPassword := helper.HashAndSalt(pwd)

		userRepoPGMock := user_repo_pg_mock.NewUserRepoPGMock()
		userRepoPGMock.On("GetByEmail", ctx, "ihsan@gmail.com").Return(domain.User{
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

		userMatcher := mock.MatchedBy(func(user *domain.User) bool {

			if user == nil {
				return false
			}

			if user.Name == "" || user.Email == "" || user.Password == "" || user.Username == "" {
				return false
			}

			compare := helper.ComparePasswords(user.Password, []byte("!H%$Nlkj123"))

			return compare
		})

		userRepoPGMock.On("RegisterUser", ctx, userMatcher).Return(nil)
		userRepoPGMock.On("GetByEmailAndUsername", ctx, "ihsan@gmail.com", "ihsanbudiman").Return(domain.User{}, nil)

		userUseCase := NewUserUseCase(userRepoPGMock)

		_, err := userUseCase.Register(ctx, domain.User{
			Email:    "ihsan@gmail.com",
			Password: "!H%$Nlkj123",
			Name:     "ihsan",
			Username: "ihsanbudiman",
		})

		if err != nil {
			t.Errorf("err %v", err)
		}

	})

}
