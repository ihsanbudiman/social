package domain

import "gorm.io/gorm"

type Handler interface {
	Run(db *gorm.DB)
}
