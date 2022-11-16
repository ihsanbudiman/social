package thread_usecase

import (
	"context"
	"errors"
	"social/domain"
	thread_repo_pg_mock "social/services/thread/repo/pg/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

func Test_threadUseCase_GetReplies(t *testing.T) {

	t.Run("success but empty", func(t *testing.T) {
		ctx := context.Background()
		threadRepoMock := thread_repo_pg_mock.NewThreadRepoPgMock()
		threadRepoMock.On("GetReplies", ctx, uint(1), 1).Return([]domain.Thread{}, nil)
		threadUseCase := NewThreadUseCase(threadRepoMock)
		threads, err := threadUseCase.GetReplies(ctx, uint(1), 1)

		assert.NoError(t, err)
		assert.Empty(t, threads)
	})

	t.Run("success with data", func(t *testing.T) {
		ctx := context.Background()
		threadRepoMock := thread_repo_pg_mock.NewThreadRepoPgMock()
		threadRepoMock.On("GetReplies", ctx, uint(1), 1).Return([]domain.Thread{
			{
				Model: gorm.Model{
					ID: 1,
				},
				Content:    "1234567890",
				ReplyTo:    null.IntFrom(1),
				ReplyCount: 1,
				LikeCount:  1,
				UserID:     1,
				User: &domain.User{
					Model: gorm.Model{
						ID: 1,
					},
				},
				Replies:      &[]domain.Thread{},
				ThreadPhotos: &[]domain.ThreadPhoto{},
			},
		}, nil)
		threadUseCase := NewThreadUseCase(threadRepoMock)
		threads, err := threadUseCase.GetReplies(ctx, uint(1), 0)

		assert.NoError(t, err)
		assert.NotEmpty(t, threads)
	})

	t.Run("threads should be empty", func(t *testing.T) {
		ctx := context.Background()

		threadRepoMock := thread_repo_pg_mock.NewThreadRepoPgMock()
		threadUseCase := NewThreadUseCase(threadRepoMock)
		threads, err := threadUseCase.GetReplies(ctx, uint(0), 0)

		assert.NoError(t, err)
		assert.Empty(t, threads)
	})

	t.Run("should be error when call getreplies on repo", func(t *testing.T) {
		ctx := context.Background()
		threadRepoMock := thread_repo_pg_mock.NewThreadRepoPgMock()
		threadRepoMock.On("GetReplies", ctx, uint(1), 1).Return([]domain.Thread{}, errors.New("terjadi kesalahan"))
		threadUseCase := NewThreadUseCase(threadRepoMock)
		threads, err := threadUseCase.GetReplies(ctx, uint(1), 1)

		assert.Error(t, err)
		assert.Empty(t, threads)
	})
}
