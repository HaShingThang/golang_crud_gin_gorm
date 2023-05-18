package helpers

import (
	"github.com/HaShingThang/golang_crud_gin_gorm/data/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseHandler(ctx *gin.Context, statusCode int, message string, data interface{}) {
	response := response.Response{
		Code:    statusCode,
		Status:  http.StatusText(statusCode),
		Message: message,
		Data:    data,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, response)
}
