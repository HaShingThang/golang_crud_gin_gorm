package interfaces

import (
	"errors"

	"github.com/HaShingThang/golang_crud_gin_gorm/data/request"
	"github.com/HaShingThang/golang_crud_gin_gorm/helpers"
	"github.com/HaShingThang/golang_crud_gin_gorm/models"
	"gorm.io/gorm"
)

type UsersInterfaceImpl struct {
	Db *gorm.DB
}

func NewUsersInterfaceImpl(Db *gorm.DB) UsersInterface {
	return &UsersInterfaceImpl{Db: Db}
}

// Delete implements UsersInterface.
func (user *UsersInterfaceImpl) Delete(userId int) {
	var users models.User
	result := user.Db.Where("id = ?", userId).Delete(&users)
	helpers.ErrorHandler(result.Error)
}

// FindAll implements UsersInterface.
func (user *UsersInterfaceImpl) FindAll() []models.User {
	var users []models.User
	result := user.Db.Preload("Posts").Find(&users)
	helpers.ErrorHandler(result.Error)
	return users
}

// FindById implements UsersInterface.
func (user *UsersInterfaceImpl) FindById(userId int) (models.User, error) {
	var users models.User
	result := user.Db.Preload("Posts").First(&users, userId)
	if result.Error != nil {
		return users, errors.New("user not found")
	}
	return users, nil
}

// FindByUsername implements UsersInterface.
func (user *UsersInterfaceImpl) FindByUsername(username string) (models.User, error) {
	var users models.User
	result := user.Db.Preload("Posts").First(&users, "username = ?", username)
	if result.Error != nil {
		return users, errors.New("invalid username or password")
	}
	return users, nil
}

// FindByEmail implements UsersInterface.
func (user *UsersInterfaceImpl) FindByEmail(email string) (models.User, error) {
	var users models.User
	result := user.Db.Preload("Posts").First(&users, "email = ?", email)
	if result.Error != nil {
		return users, errors.New("invalid email or password")
	}
	return users, nil
}

// Save implements UsersInterface.
func (user *UsersInterfaceImpl) Save(users models.User) {
	result := user.Db.Create(&users)
	helpers.ErrorHandler(result.Error)
}

// Update implements UsersInterface.
func (user *UsersInterfaceImpl) Update(users models.User) {
	var updateUsers = request.UpdateUserRequest{
		Id:       users.Id,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}

	result := user.Db.Model(&users).Updates(updateUsers)
	helpers.ErrorHandler(result.Error)
}
