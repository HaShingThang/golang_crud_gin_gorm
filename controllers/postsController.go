package controllers

import (
	"net/http"
	"strconv"

	"github.com/HaShingThang/golang_crud_gin_gorm/data/request"
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
	if err != nil {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Invalid post create request.", nil)
		return
	}

	//Required Title
	if createPostsRequest.Title == "" {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Title is required.", nil)
		return
	}
	//Required Description
	if createPostsRequest.Description == "" {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Description is required.", nil)
		return
	}

	controller.PostsService.Create(createPostsRequest)
	helpers.ResponseHandler(ctx, http.StatusOK, "Created Post Success.", nil)
}

// Update Post
func (controller *PostsController) Update(ctx *gin.Context) {
	updatePostsRequest := request.UpdatePostsRequest{}
	err := ctx.ShouldBindJSON(&updatePostsRequest)
	if err != nil {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Invalid post update request.", nil)
		return
	}

	postId := ctx.Param("postId")
	id, err := strconv.Atoi(postId)
	if err != nil {
		// PostId must be integer
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Invalid postId.", nil)
		return
	}

	postRes := controller.PostsService.FindById(id)
	if postRes.Id == 0 {
		//Post is not found
		helpers.ResponseHandler(ctx, http.StatusNotFound, "Post not found.", nil)
		return
	}
	updatePostsRequest.Id = id
	updatedPost := controller.PostsService.Update(updatePostsRequest)
	helpers.ResponseHandler(ctx, http.StatusOK, "Updated Post Success.", updatedPost)
}

// Delete post
func (controller *PostsController) Delete(ctx *gin.Context) {
	postId := ctx.Param("postId")
	id, err := strconv.Atoi(postId)
	if err != nil {
		// PostId must be integer
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Invalid postId.", nil)
		return
	}
	postRes := controller.PostsService.FindById(id)
	if postRes.Id == 0 {
		//Post is not found
		helpers.ResponseHandler(ctx, http.StatusNotFound, "Post not found.", nil)
		return
	}
	controller.PostsService.Delete(id)
	helpers.ResponseHandler(ctx, http.StatusOK, "Deleted Post Success.", nil)
}

// Get Post By Id
func (controller *PostsController) FindById(ctx *gin.Context) {
	postId := ctx.Param("postId")
	id, err := strconv.Atoi(postId)
	if err != nil {
		// PostId must be integer
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Invalid postId.", nil)
		return
	}

	postRes := controller.PostsService.FindById(id)
	if postRes.Id == 0 {
		//Post is not found
		helpers.ResponseHandler(ctx, http.StatusNotFound, "Post not found.", nil)
		return
	}

	helpers.ResponseHandler(ctx, http.StatusOK, "Get Post Success.", postRes)
}

// Get All Posts
func (controller *PostsController) FindAll(ctx *gin.Context) {
	postRes := controller.PostsService.FindAll()
	if len(postRes) == 0 {
		// not found posts
		helpers.ResponseHandler(ctx, http.StatusNotFound, "No posts found.", postRes)
		return
	}
	helpers.ResponseHandler(ctx, http.StatusOK, "Get All Post Success.", postRes)
}
