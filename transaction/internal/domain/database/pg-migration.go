package db

import (
	"github.com/jinzhu/gorm"
	models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Transaction{})
}
