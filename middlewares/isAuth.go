package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/HaShingThang/golang_crud_gin_gorm/helpers"
	"github.com/HaShingThang/golang_crud_gin_gorm/interfaces"
	"github.com/HaShingThang/golang_crud_gin_gorm/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func IsAuth(usersInterface interfaces.UsersInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizeHeader := ctx.GetHeader("Authorization")
		token := strings.TrimPrefix(authorizeHeader, "Bearer ")

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": "You are not logged in."})
			return
		}

		err := godotenv.Load(".env")
		helpers.ErrorHandler(err)
		tokenSecret := os.Getenv("TOKEN_SECRET")
		sub, err := utils.ValidateToken(token, tokenSecret)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": err.Error()})
			return
		}

		id, err_id := strconv.Atoi(fmt.Sprint(sub))
		helpers.ErrorHandler(err_id)
		result, err := usersInterface.FindById(id)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": "User not found."})
			return
		}

		ctx.Set("currentUser", result.Username)
		ctx.Next()
	}
}
