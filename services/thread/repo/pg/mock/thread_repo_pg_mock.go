package thread_repo_pg_mock

import (
	"context"
	"social/domain"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type ThreadRepoPgMock struct {
	mock.Mock
}

// GetReplies implements domain.ThreadRepo
func (t ThreadRepoPgMock) GetReplies(ctx context.Context, threadID uint, page int) ([]domain.Thread, error) {
	args := t.Called(ctx, threadID, page)
	return args.Get(0).([]domain.Thread), args.Error(1)
}

// GetThread implements domain.ThreadRepo
func (t ThreadRepoPgMock) GetThread(ctx context.Context, threadID uint) (domain.Thread, error) {
	args := t.Called(ctx, threadID)
	return args.Get(0).(domain.Thread), args.Error(1)
}

// Insert implements domain.ThreadRepo
func (t ThreadRepoPgMock) Insert(ctx context.Context, thread *domain.Thread) error {
	args := t.Called(ctx, thread)
	return args.Error(0)
}

// InsertThreadPhoto implements domain.ThreadRepo
func (t ThreadRepoPgMock) InsertThreadPhoto(ctx context.Context, threadPhoto *[]domain.ThreadPhoto) error {
	args := t.Called(ctx, threadPhoto)
	return args.Error(0)
}

// WithTx implements domain.ThreadRepo
func (t ThreadRepoPgMock) WithTx(ctx context.Context, tx *gorm.DB) domain.ThreadRepo {
	return &ThreadRepoPgMock{}
}

func NewThreadRepoPgMock() ThreadRepoPgMock {
	return ThreadRepoPgMock{}
}
