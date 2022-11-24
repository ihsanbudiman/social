package config

import (
	"social/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=ihsan password=ihsan dbname=social port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Exec("select 1").Error
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.User{}, &domain.Thread{}, &domain.ThreadPhoto{}, &domain.Like{}, &domain.Kota{}, &domain.Jadwal{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
