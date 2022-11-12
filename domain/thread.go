package domain

import "gorm.io/gorm"

type Thread struct {
	gorm.Model
	Content      string        `json:"content" gorm:"type:text"`
	ThreadPhotos []ThreadPhoto `json:"thread_photos" gorm:"foreignKey:ThreadID"`
	IsReply      bool          `json:"is_reply" gorm:"default:false"`
	ReplyTo      uint          `json:"reply_to" gorm:"default:0,index:idx_reply_to"`
	ReplyCount   uint          `json:"reply_count" gorm:"default:0"`
	LikeCount    uint          `json:"like_count" gorm:"default:0"`
	IsLiked      bool          `json:"is_liked" gorm:"-"`
	UserID       uint          `json:"user_id" gorm:"index:idx_user_id"`
	User         User          `json:"user" gorm:"foreignKey:UserID"`
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
