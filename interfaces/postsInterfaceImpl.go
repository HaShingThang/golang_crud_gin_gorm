package interfaces

import (
	"errors"

	"github.com/HaShingThang/golang_crud_gin_gorm/data/request"
	"github.com/HaShingThang/golang_crud_gin_gorm/helpers"
	"github.com/HaShingThang/golang_crud_gin_gorm/models"
	"gorm.io/gorm"
)

type PostsInterfaceImpl struct {
	Db *gorm.DB
}

func NewPostsInterfaceImpl(Db *gorm.DB) PostsInterface {
	return &PostsInterfaceImpl{Db: Db}
}

// Create implements PostsInterface.
func (post *PostsInterfaceImpl) Save(posts models.Posts) {
	result := post.Db.Create(&posts)
	helpers.ErrorHandler(result.Error)
}

// Delete implements PostsInterface.
func (post *PostsInterfaceImpl) Delete(postsId int) {
	var posts models.Posts
	result := post.Db.Where("id = ?", postsId).Delete(&posts)
	helpers.ErrorHandler(result.Error)
}

// FindAll implements PostsInterface.
func (post *PostsInterfaceImpl) FindAll() []models.Posts {
	var posts []models.Posts
	result := post.Db.Find(&posts)
	helpers.ErrorHandler(result.Error)
	return posts
}

// FindById implements PostsInterface.
func (post *PostsInterfaceImpl) FindById(postsId int) (posts models.Posts, err error) {
	var pst models.Posts
	result := post.Db.Find(&pst, postsId)
	if result != nil {
		return pst, nil
	} else {
		return pst, errors.New("tag is not found")
	}
}

// Update implements PostsInterface.
func (post *PostsInterfaceImpl) Update(posts models.Posts) {
	var updatepost = request.UpdatePostsRequest{
		Id:          posts.Id,
		Title:       posts.Title,
		Description: posts.Description,
	}
	result := post.Db.Model(&posts).Updates(updatepost)
	helpers.ErrorHandler(result.Error)
}
