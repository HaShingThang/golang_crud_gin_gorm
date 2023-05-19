package services

import (
	"github.com/HaShingThang/golang_crud_gin_gorm/data/request"
	"github.com/HaShingThang/golang_crud_gin_gorm/models"
)

type Authservice interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUserRequest)
	FindByEmail(email string) models.User
}
