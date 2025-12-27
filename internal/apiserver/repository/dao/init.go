package dao

import (
	"github.com/jwcen/mars/internal/apiserver/repository/model"
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(&model.UserM{})
}
