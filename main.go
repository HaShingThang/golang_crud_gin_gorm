package main

import (
	"fmt"
	"github.com/HaShingThang/golang_crud_gin_gorm/config"
	"github.com/HaShingThang/golang_crud_gin_gorm/controllers"
	"github.com/HaShingThang/golang_crud_gin_gorm/interfaces"
	"github.com/HaShingThang/golang_crud_gin_gorm/migrations"
	"github.com/HaShingThang/golang_crud_gin_gorm/routers"
	"github.com/HaShingThang/golang_crud_gin_gorm/services"
	"github.com/go-playground/validator/v10"
)

func main() {

	// Database
	db, _ := config.ConnectDB()
	validate := validator.New()

	//Table
	migrations.RunMigrations(db)

	//Users
	usersInterface := interfaces.NewUsersInterfaceImpl(db)
	authService := services.NewAuthServiceImpl(usersInterface, validate)
	authController := controllers.NewAuthController(authService)
	usersController := controllers.NewUsercontroller(usersInterface)

	//Post
	postsInterface := interfaces.NewPostsInterfaceImpl(db)
	postsService := services.NewPostsServiceImpl(postsInterface, validate)
	postsController := controllers.NewPostsController(postsService)

	//Router
	routes := routers.Router(authController, usersController, usersInterface, postsController)

	fmt.Println("Server running on http://localhost:8080")
	routes.Run(":8080")
}
