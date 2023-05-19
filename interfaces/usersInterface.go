package interfaces

import "github.com/HaShingThang/golang_crud_gin_gorm/models"

type UsersInterface interface {
	Save(users models.User)
	Update(users models.User)
	Delete(userId int)
	FindById(userId int) (models.User, error)
	FindAll() []models.User
	FindByEmail(email string) (models.User, error)
	FindByUsername(username string) (models.User, error)
}
