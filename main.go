package main

import (
	"fmt"
	"net/http"

	"github.com/HaShingThang/golang_crud_gin_gorm/config"
	"github.com/HaShingThang/golang_crud_gin_gorm/controllers"
	"github.com/HaShingThang/golang_crud_gin_gorm/helpers"
	"github.com/HaShingThang/golang_crud_gin_gorm/interfaces"
	"github.com/HaShingThang/golang_crud_gin_gorm/models"
	"github.com/HaShingThang/golang_crud_gin_gorm/routers"
	"github.com/HaShingThang/golang_crud_gin_gorm/services"
	"github.com/go-playground/validator/v10"
)

func main() {

	// Database
	db := config.ConnectDB()
	validate := validator.New()

	//Table
	db.Table("users").AutoMigrate(&models.Users{})
	db.Table("posts").AutoMigrate(&models.Posts{})

	//Users
	usersInterface := interfaces.NewUsersInterfaceImpl(db)
	authService := services.NewAuthServiceImpl(usersInterface, validate)
	authController := controllers.NewAuthController(authService)
	usersController := controllers.NewUsercontroller(usersInterface)

	//Post
	db.Table("posts").AutoMigrate(&models.Posts{})
	postsInterface := interfaces.NewPostsInterfaceImpl(db)
	postsService := services.NewPostsServiceImpl(postsInterface, validate)
	postsController := controllers.NewpostsController(postsService)

	//Router
	// routes := routers.NewUsersRouter(usersInterface, authController, usersController)
	// routes := routers.NewPostsRouter(postsController)
	routes := routers.Router(authController, usersController, usersInterface, postsController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	fmt.Println("Server running on http://localhost:8080")

	err := server.ListenAndServe()
	helpers.ErrorHandler(err)
}
