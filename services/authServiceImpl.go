package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/HaShingThang/golang_crud_gin_gorm/data/request"
	"github.com/HaShingThang/golang_crud_gin_gorm/helpers"
	"github.com/HaShingThang/golang_crud_gin_gorm/interfaces"
	"github.com/HaShingThang/golang_crud_gin_gorm/models"
	"github.com/HaShingThang/golang_crud_gin_gorm/utils"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type AuthServiceImpl struct {
	UsersInterface interfaces.UsersInterface
	Validate       *validator.Validate
}

func NewAuthServiceImpl(usersInterface interfaces.UsersInterface, validate *validator.Validate) Authservice {
	return &AuthServiceImpl{
		UsersInterface: usersInterface,
		Validate:       validate,
	}
}

// Register implements AuthService.
func (auth *AuthServiceImpl) Register(users request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(users.Password)
	helpers.ErrorHandler(err)

	newUser := models.User{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}
	auth.UsersInterface.Save(newUser)
}

// Login implements Authservice.
func (auth *AuthServiceImpl) Login(users request.LoginRequest) (string, error) {
	err := godotenv.Load(".env")
	helpers.ErrorHandler(err)
	tokenExpireInStr := os.Getenv("TOKEN_EXPIRED_IN")
	tokenSecret := os.Getenv("TOKEN_SECRET")

	tokenDuration, err := time.ParseDuration(tokenExpireInStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse TOKEN_EXPIRED_IN: %w", err)
	}

	newUser, err := auth.UsersInterface.FindByEmail(users.Email)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	verifyPassword := utils.VerifyPassword(newUser.Password, users.Password)
	if verifyPassword != nil {
		return "", errors.New("invalid email or password")
	}
	token, err_token := utils.GenerateToken(tokenDuration, newUser.Id, tokenSecret)
	helpers.ErrorHandler(err_token)
	return token, nil
}

// FindByEmail implements Authservice.
func (auth *AuthServiceImpl) FindByEmail(email string) models.User {
	user, _ := auth.UsersInterface.FindByEmail(email)

	return user
}
