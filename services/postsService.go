package services

import (
	"github.com/HaShingThang/golang_crud_gin_gorm/data/request"
	"github.com/HaShingThang/golang_crud_gin_gorm/data/response"
)

type PostsService interface {
	Create(posts request.CreatePostRequest) response.PostResponse
	Update(posts request.UpdatePostRequest) response.PostResponse
	Delete(postsId int)
	FindById(postsId int) response.PostResponse
	FindAll() []response.PostResponse
}
