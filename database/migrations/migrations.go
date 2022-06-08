package migrations

import (
	"github.com/FredsZDev/WebApi_GO_B2/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Game{})

}
