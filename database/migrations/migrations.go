package migrations

import (
	"github.com/flaviokicis/go-restapi/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
