package routers

import (
	"net/http"

	"github.com/HaShingThang/golang_crud_gin_gorm/controllers"
	"github.com/gin-gonic/gin"
)

func NewPostsRouter(postsController *controllers.PostsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Default Router is OK!")
	})
	baseRouter := router.Group("/api")
	postsRouter := baseRouter.Group("/posts")
	postsRouter.GET("", postsController.FindAll)
	postsRouter.GET("/:postId", postsController.FindById)
	postsRouter.POST("", postsController.Create)
	postsRouter.PATCH("/:postId", postsController.Update)
	postsRouter.DELETE("/:postId", postsController.Delete)

	return router
}
