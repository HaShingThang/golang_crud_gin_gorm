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

func NewPostsController(service services.PostsService) *PostsController {
	return &PostsController{
		PostsService: service,
	}
}

// Create post
func (controller *PostsController) Create(ctx *gin.Context) {
	createPostsRequest := request.CreatePostRequest{}
	err := ctx.ShouldBindJSON(&createPostsRequest)
	if err != nil {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Invalid post create.", nil)
		return
	}

	// Required Title
	if createPostsRequest.Title == "" {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Title is required.", nil)
		return
	}
	// Required Description
	if createPostsRequest.Description == "" {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Description is required.", nil)
		return
	}

	post := controller.PostsService.Create(createPostsRequest)
	helpers.ResponseHandler(ctx, http.StatusOK, "Created Post Success.", post)
}

// Update Post
func (controller *PostsController) Update(ctx *gin.Context) {
	updatePostsRequest := request.UpdatePostRequest{}
	err := ctx.ShouldBindJSON(&updatePostsRequest)
	if err != nil {
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Invalid post update request.", nil)
		return
	}

	postId := ctx.Param("postId")
	id, err := strconv.Atoi(postId)
	if err != nil {
		// PostId must be an integer
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Invalid postId.", nil)
		return
	}

	postRes := controller.PostsService.FindById(id)
	if postRes.Id == 0 {
		// Post is not found
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
		// PostId must be an integer
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Invalid postId.", nil)
		return
	}

	postRes := controller.PostsService.FindById(id)
	if postRes.Id == 0 {
		// Post is not found
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
		// PostId must be an integer
		helpers.ResponseHandler(ctx, http.StatusBadRequest, "Invalid postId.", nil)
		return
	}

	postRes := controller.PostsService.FindById(id)
	if postRes.Id == 0 {
		// Post is not found
		helpers.ResponseHandler(ctx, http.StatusNotFound, "Post not found.", nil)
		return
	}
	// // Fetch user associated with the post
	// userRes := controller.PostsService.GetUserByPostId(id)
	// postRes.UserId = userRes.Id

	helpers.ResponseHandler(ctx, http.StatusOK, "Get Post Success.", postRes)
}

// Get All Posts
func (controller *PostsController) FindAll(ctx *gin.Context) {
	postRes := controller.PostsService.FindAll()
	if len(postRes) == 0 {
		// No posts found
		helpers.ResponseHandler(ctx, http.StatusNotFound, "No posts found.", postRes)
		return
	}

	// Fetch users associated with the posts
	// for i, post := range postRes {
	// 	userRes := controller.PostsService.GetUserByPostId(post.Id)
	// 	postRes[i].User = userRes
	// }

	helpers.ResponseHandler(ctx, http.StatusOK, "Get All Post Success.", postRes)
}
