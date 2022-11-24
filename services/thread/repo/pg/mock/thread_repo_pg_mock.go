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

// CheckLiked implements domain.ThreadRepo
func (t ThreadRepoPgMock) CheckLiked(ctx context.Context, threadID uint, userID uint) (bool, error) {
	args := t.Called(ctx, threadID, userID)
	return args.Bool(0), args.Error(1)
}

// LikeThread implements domain.ThreadRepo
func (t ThreadRepoPgMock) LikeThread(ctx context.Context, threadID uint, userID uint) (domain.Like, error) {
	args := t.Called(ctx, threadID, userID)
	return args.Get(0).(domain.Like), args.Error(1)
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
func (t ThreadRepoPgMock) InsertThreadPhoto(ctx context.Context, threadPhoto *domain.ThreadPhoto) error {
	args := t.Called(ctx, threadPhoto)
	return args.Error(0)
}

func (t ThreadRepoPgMock) UnlikeThread(ctx context.Context, threadID, userID uint) error {
	args := t.Called(ctx, threadID, userID)
	return args.Error(0)
}

// WithTx implements domain.ThreadRepo
func (t ThreadRepoPgMock) WithTx(ctx context.Context, tx *gorm.DB) domain.ThreadRepo {
	return &ThreadRepoPgMock{
		Mock: mock.Mock{},
	}
}

func NewThreadRepoPgMock() ThreadRepoPgMock {
	return ThreadRepoPgMock{}
}
