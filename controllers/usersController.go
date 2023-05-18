package controllers

import (
	"net/http"

	"github.com/HaShingThang/golang_crud_gin_gorm/helpers"
	"github.com/HaShingThang/golang_crud_gin_gorm/interfaces"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	UsersInterface interfaces.UsersInterface
}

func NewUsercontroller(interfaces interfaces.UsersInterface) *UsersController {
	return &UsersController{UsersInterface: interfaces}
}

func (controller *UsersController) GetUsers(ctx *gin.Context) {
	users := controller.UsersInterface.FindAll()
	helpers.ResponseHandler(ctx, http.StatusOK, "Get All Users Success.", users)
}
