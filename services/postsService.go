package services

import (
	"github.com/HaShingThang/golang_crud_gin_gorm/data/request"
	"github.com/HaShingThang/golang_crud_gin_gorm/data/response"
)

type PostsService interface {
	Create(posts request.CreatePostsRequest) response.PostsResponse
	Update(posts request.UpdatePostsRequest) response.PostsResponse
	Delete(postsId int)
	FindById(postsId int) response.PostsResponse
	FindAll() []response.PostsResponse
}
