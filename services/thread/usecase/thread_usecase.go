package thread_usecase

import (
	"context"
	"errors"
	"social/domain"

	"gorm.io/gorm"
)

type threadUseCase struct {
	repo domain.ThreadRepo
}

// GetReplies implements domain.ThreadUseCase
func (t threadUseCase) GetReplies(ctx context.Context, threadID uint, page int) ([]domain.Thread, error) {
	if threadID == 0 {
		return []domain.Thread{}, nil
	}

	if page <= 0 {
		page = 1
	}

	threads, err := t.repo.GetReplies(ctx, threadID, page)
	if err != nil {
		return threads, err
	}

	return threads, nil
}

// GetThread implements domain.ThreadUseCase
func (t threadUseCase) GetThread(ctx context.Context, threadID uint) (domain.Thread, error) {
	if threadID == 0 {
		return domain.Thread{}, nil
	}

	thread, err := t.repo.GetThread(ctx, threadID)
	if err != nil {
		return domain.Thread{}, err
	}

	return thread, nil
}

// CreateThread implements domain.ThreadUseCase
func (t threadUseCase) CreateThread(ctx context.Context, thread domain.Thread) (domain.Thread, error) {
	// validation thread
	if thread.UserID == 0 {
		return domain.Thread{}, errors.New("user id is required")
	}

	if thread.Content == "" {
		return domain.Thread{}, errors.New("content is required")
	}

	// insert thread
	err := t.repo.Insert(ctx, &thread)
	if err != nil {
		return thread, err
	}

	return thread, nil
}

// WithTx implements domain.ThreadUseCase
func (t threadUseCase) WithTx(ctx context.Context, tx *gorm.DB) domain.ThreadUseCase {
	return &threadUseCase{
		repo: t.repo.WithTx(ctx, tx),
	}
}

func NewThreadUseCase(repo domain.ThreadRepo) domain.ThreadUseCase {
	return &threadUseCase{
		repo: repo,
	}
}
