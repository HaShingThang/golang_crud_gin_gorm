package interfaces

import "github.com/HaShingThang/golang_crud_gin_gorm/models"

type PostsInterface interface {
	Save(posts models.Posts)
	Update(posts models.Posts)
	Delete(postsId int)
	FindById(postsId int) (posts models.Posts, err error)
	FindAll() []models.Posts
}
