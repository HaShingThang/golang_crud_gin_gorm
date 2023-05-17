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

	db.Table("posts").AutoMigrate(&models.Posts{})

	// Interface
	postsInterface := interfaces.NewPostsInterfaceImpl(db)

	//Service
	postsService := services.NewPostsServiceImpl(postsInterface, validate)

	//Controller
	postController := controllers.NewpostsController(postsService)

	//Router
	routes := routers.NewPostsRouter(postController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	fmt.Println("Server running on http://localhost:8080")

	err := server.ListenAndServe()
	helpers.ErrorHandler(err)
}
