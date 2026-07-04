package repository

import (
	"github.com/xurceo/plain-tcms/db"
	"gorm.io/gorm"
)

func getDB() *gorm.DB {
	return db.DB
}
