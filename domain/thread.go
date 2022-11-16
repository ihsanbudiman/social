package domain

import (
	"context"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	Content      string         `json:"content" gorm:"type:text"`
	ReplyTo      null.Int       `json:"reply_to" gorm:"index:reply_to_idx;"`
	IsReply      uint           `json:"is_reply" gorm:"default:0"`
	ReplyCount   uint           `json:"reply_count" gorm:"default:0"`
	LikeCount    uint           `json:"like_count" gorm:"default:0"`
	UserID       uint           `json:"user_id" gorm:"index:idx_user_id"`
	User         *User          `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT;"`
	Replies      *[]Thread      `json:"replies" gorm:"foreignKey:ReplyTo;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT;"`
	ThreadPhotos *[]ThreadPhoto `json:"thread_photos" gorm:"foreignKey:ThreadID;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT;"`
}

type ThreadPhoto struct {
	gorm.Model
	ThreadID uint   `gorm:"not null,index:idx_thread_id" json:"thread_id"`
	FileUrl  string `gorm:"not null" json:"file_url"`
}

type Like struct {
	gorm.Model
	UserID   uint `gorm:"not null,index:idx_user_id" json:"user_id"`
	ThreadID uint `gorm:"not null,index:idx_thread_id" json:"thread_id"`
}

type ThreadRepo interface {
	Insert(ctx context.Context, thread *Thread) error
	InsertThreadPhoto(ctx context.Context, threadPhoto *ThreadPhoto) error
	GetThread(ctx context.Context, threadID uint) (Thread, error)
	GetReplies(ctx context.Context, threadID uint, page int) ([]Thread, error)
	WithTx(ctx context.Context, tx *gorm.DB) ThreadRepo
}

type ThreadUseCase interface {
	CreateThread(ctx context.Context, thread Thread) (Thread, error)
	InsertThreadPhoto(ctx context.Context, threadPhoto *ThreadPhoto) error
	GetThread(ctx context.Context, threadID uint) (Thread, error)
	GetReplies(ctx context.Context, threadID uint, page int) ([]Thread, error)
	WithTx(ctx context.Context, tx *gorm.DB) ThreadUseCase
}
