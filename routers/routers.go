package routers

import (
	"net/http"

	"github.com/HaShingThang/golang_crud_gin_gorm/controllers"
	"github.com/HaShingThang/golang_crud_gin_gorm/interfaces"
	"github.com/HaShingThang/golang_crud_gin_gorm/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(authController *controllers.AuthController, usersController *controllers.UsersController, usersInterface interfaces.UsersInterface, postsController *controllers.PostsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Default Router is OK!")
	})

	apiRouter := router.Group("/api")
	AuthRouter(apiRouter, authController)
	UsersRouter(apiRouter, usersInterface, usersController)
	PostsRouter(apiRouter, postsController, usersInterface)
	return router
}

// Auth Router
func AuthRouter(router *gin.RouterGroup, authController *controllers.AuthController) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", authController.Register)
		authRouter.POST("/login", authController.Login)
	}
}

// User Router
func UsersRouter(router *gin.RouterGroup, usersInterface interfaces.UsersInterface, usersController *controllers.UsersController) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("", middlewares.IsAuth(usersInterface), usersController.GetUsers)
	}
}

// Post Router
func PostsRouter(router *gin.RouterGroup, postsController *controllers.PostsController, usersInterface interfaces.UsersInterface) {
	postsRouter := router.Group("/posts")
	{
		postsRouter.GET("", middlewares.IsAuth(usersInterface), postsController.FindAll)
		postsRouter.GET("/:postId", middlewares.IsAuth(usersInterface), postsController.FindById)
		postsRouter.POST("", middlewares.IsAuth(usersInterface), postsController.Create)
		postsRouter.PATCH("/:postId", middlewares.IsAuth(usersInterface), postsController.Update)
		postsRouter.DELETE("/:postId", middlewares.IsAuth(usersInterface), postsController.Delete)
	}
}
