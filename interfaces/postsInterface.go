package interfaces

import "github.com/HaShingThang/golang_crud_gin_gorm/models"

type PostsInterface interface {
	Save(posts models.Post)
	Update(posts models.Post)
	Delete(postsId int)
	FindById(postsId int) (posts models.Post, err error)
	FindAll() []models.Post
}
