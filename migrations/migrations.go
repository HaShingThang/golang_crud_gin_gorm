package migrations

import (
	"github.com/HaShingThang/golang_crud_gin_gorm/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
}
