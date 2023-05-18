package interfaces

import "github.com/HaShingThang/golang_crud_gin_gorm/models"

type UsersInterface interface {
	Save(users models.Users)
	Update(users models.Users)
	Delete(userId int)
	FindById(userId int) (models.Users, error)
	FindAll() []models.Users
	FindByEmail(email string) (models.Users, error)
	FindByUsername(username string) (models.Users, error)
}
