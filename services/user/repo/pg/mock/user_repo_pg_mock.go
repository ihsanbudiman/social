package user_repo_pg_mock

import (
	"context"
	"social/domain"

	"github.com/stretchr/testify/mock"
)

type userRepoPGMock struct {
	mock.Mock
}

func (m userRepoPGMock) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	args := m.Called(ctx, email)

	return args.Get(0).(domain.User), args.Error(1)
}

func (m userRepoPGMock) UpdateUser(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)

	return args.Error(1)
}

func (m userRepoPGMock) GetByID(ctx context.Context, ID uint) (domain.User, error) {
	user := domain.User{}
	args := m.Called(ctx, ID, user)

	return args.Get(2).(domain.User), args.Error(1)
}

func (m userRepoPGMock) RegisterUser(ctx context.Context, user *domain.User) error {

	args := m.Called(ctx, user)

	return args.Error(0)
}

func (m userRepoPGMock) GetByEmailAndUsername(ctx context.Context, email string, username string) (domain.User, error) {
	args := m.Called(ctx, email, username)

	return args.Get(0).(domain.User), args.Error(1)
}

func NewUserRepoPGMock() userRepoPGMock {
	return userRepoPGMock{}
}
