package repository

import "gorm.io/gorm"

type GORMRepository struct {
	db *gorm.DB
}
