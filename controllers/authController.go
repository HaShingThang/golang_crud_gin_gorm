package controllers

import (
	"net/http"

	"github.com/HaShingThang/golang_crud_gin_gorm/data/request"
	"github.com/HaShingThang/golang_crud_gin_gorm/data/response"
	"github.com/HaShingThang/golang_crud_gin_gorm/helpers"
	"github.com/HaShingThang/golang_crud_gin_gorm/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService services.Authservice
}

func NewAuthController(service services.Authservice) *AuthController {
	return &AuthController{
		AuthService: service,
	}
}

// Register Controller
func (controller *AuthController) Register(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helpers.ErrorHandler(err)
	if err != nil {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Invalid user create.", nil)
		return
	}
	// Check if email already exists
	existingUser := controller.AuthService.FindByEmail(createUserRequest.Email)
	if existingUser.Id != 0 {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Email already exists.", nil)
		return
	}

	//Required Username
	if createUserRequest.Username == "" {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Username is required.", nil)
		return
	}
	//Required Email
	if createUserRequest.Email == "" {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Email is required.", nil)
		return
	}
	//Required Password
	if createUserRequest.Password == "" || len(createUserRequest.Password) < 6 {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Password must be at least 6 chars long.", nil)
		return
	}

	controller.AuthService.Register(createUserRequest)
	helpers.ResponseHandler(ctx, http.StatusOK, "Created User Success.", nil)
}

// Login Controller
func (controller *AuthController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helpers.ErrorHandler(err)
	token, err_token := controller.AuthService.Login(loginRequest)
	if err_token != nil {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Invalid email or password.", nil)
		return
	}
	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}
	helpers.ResponseHandler(ctx, http.StatusOK, "Login Success.", resp)
}
