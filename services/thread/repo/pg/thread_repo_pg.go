package thread_repo_pg

import (
	"context"
	"social/domain"

	"gorm.io/gorm"
)

type threadRepoPg struct {
	// gorm connection
	db *gorm.DB
}

// GetReplies implements domain.ThreadRepo
func (t threadRepoPg) GetReplies(ctx context.Context, threadID uint, page int) ([]domain.Thread, error) {
	var threads []domain.Thread

	err := t.db.Where("thread_id = ?", threadID).Limit(10).Offset((page - 1) * 10).Find(&threads).Error
	if err != nil {
		return threads, err
	}

	return threads, nil
}

// GetThread implements domain.ThreadRepo
func (t threadRepoPg) GetThread(ctx context.Context, threadID uint) (domain.Thread, error) {
	var thread domain.Thread
	err := t.db.
		Preload("ThreadPhotos").
		First(&thread, threadID).Error

	if err == gorm.ErrRecordNotFound {
		return thread, nil
	}

	if err != nil {
		return thread, err
	}

	return thread, nil
}

// WithTx implements domain.ThreadRepo
func (t threadRepoPg) WithTx(ctx context.Context, tx *gorm.DB) domain.ThreadRepo {
	return &threadRepoPg{
		db: tx,
	}
}

// Insert implements domain.ThreadRepo
func (t threadRepoPg) Insert(ctx context.Context, thread *domain.Thread) error {
	err := t.db.Create(thread).Error
	if err != nil {
		return err
	}

	return nil
}

// InsertThreadPhoto implements domain.ThreadRepo
func (t threadRepoPg) InsertThreadPhoto(ctx context.Context, threadPhoto *[]domain.ThreadPhoto) error {
	err := t.db.Create(threadPhoto).Error
	if err != nil {
		return err
	}

	return nil
}

func NewThreadRepoPg(db *gorm.DB) domain.ThreadRepo {
	return &threadRepoPg{
		db: db,
	}
}
