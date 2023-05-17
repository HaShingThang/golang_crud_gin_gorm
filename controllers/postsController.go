package controllers

import (
	"net/http"
	"strconv"

	"github.com/HaShingThang/golang_crud_gin_gorm/data/request"
	"github.com/HaShingThang/golang_crud_gin_gorm/data/response"
	"github.com/HaShingThang/golang_crud_gin_gorm/helpers"
	"github.com/HaShingThang/golang_crud_gin_gorm/services"
	"github.com/gin-gonic/gin"
)

type PostsController struct {
	PostsService services.PostsService
}

func NewpostsController(service services.PostsService) *PostsController {
	return &PostsController{
		PostsService: service,
	}
}

// Create post
func (controller *PostsController) Create(ctx *gin.Context) {
	createPostsRequest := request.CreatePostsRequest{}
	err := ctx.ShouldBindJSON(&createPostsRequest)
	helpers.ErrorHandler(err)

	controller.PostsService.Create(createPostsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Message: "Created Post Success",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Update post
func (controller *PostsController) Update(ctx *gin.Context) {
	updatepostsRequest := request.UpdatePostsRequest{}
	err := ctx.ShouldBindJSON(&updatepostsRequest)
	helpers.ErrorHandler(err)

	postId := ctx.Param("postId")
	id, err := strconv.Atoi(postId)
	helpers.ErrorHandler(err)
	updatepostsRequest.Id = id

	controller.PostsService.Update(updatepostsRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Message: "Updated Post Success",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete post
func (controller *PostsController) Delete(ctx *gin.Context) {
	postId := ctx.Param("postId")
	id, err := strconv.Atoi(postId)
	helpers.ErrorHandler(err)
	controller.PostsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Message: "Deleted Post Success",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Get post By Id
func (controller *PostsController) FindById(ctx *gin.Context) {
	postId := ctx.Param("postId")
	id, err := strconv.Atoi(postId)
	helpers.ErrorHandler(err)

	postRes := controller.PostsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Message: "Get Post Success",
		Data:   postRes,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Get All posts
func (controller *PostsController) FindAll(ctx *gin.Context) {
	postRes := controller.PostsService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Message: "Get All Post Success",
		Data:   postRes,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
