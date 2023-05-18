package services

import (
	"github.com/HaShingThang/golang_crud_gin_gorm/data/request"
	"github.com/HaShingThang/golang_crud_gin_gorm/data/response"
	"github.com/HaShingThang/golang_crud_gin_gorm/helpers"
	"github.com/HaShingThang/golang_crud_gin_gorm/interfaces"
	"github.com/HaShingThang/golang_crud_gin_gorm/models"
	"github.com/go-playground/validator/v10"
)

type PostsServiceImpl struct {
	PostsInterface interfaces.PostsInterface
	Validate       *validator.Validate
}

func NewPostsServiceImpl(postInterface interfaces.PostsInterface, validate *validator.Validate) PostsService {
	return &PostsServiceImpl{
		PostsInterface: postInterface,
		Validate:       validate,
	}
}

// Create implements PostsService.
func (pst *PostsServiceImpl) Create(posts request.CreatePostsRequest) response.PostsResponse {
	err := pst.Validate.Struct(posts)
	helpers.ErrorHandler(err)
	Posts := models.Posts{
		Title:       posts.Title,
		Description: posts.Description,
	}
	pst.PostsInterface.Save(Posts)
	return response.PostsResponse(Posts)
}

// Delete implements PostsService.
func (pst *PostsServiceImpl) Delete(postsId int) {
	pst.PostsInterface.Delete(postsId)
}

// FindAll implements PostsService.
func (pst *PostsServiceImpl) FindAll() []response.PostsResponse {
	result := pst.PostsInterface.FindAll()

	var posts []response.PostsResponse
	for _, value := range result {
		post := response.PostsResponse{
			Id:          value.Id,
			Title:       value.Title,
			Description: value.Description,
		}
		posts = append(posts, post)
	}

	return posts
}

// FindById implements PostsService.
func (pst *PostsServiceImpl) FindById(postsId int) response.PostsResponse {
	postData, err := pst.PostsInterface.FindById(postsId)
	helpers.ErrorHandler(err)

	postRes := response.PostsResponse{
		Id:          postData.Id,
		Title:       postData.Title,
		Description: postData.Description,
	}
	return postRes
}

// Update implements PostsService.
func (pst *PostsServiceImpl) Update(posts request.UpdatePostsRequest) response.PostsResponse {
	postData, err := pst.PostsInterface.FindById(posts.Id)
	helpers.ErrorHandler(err)
	postData.Title = posts.Title
	postData.Description = posts.Description
	pst.PostsInterface.Update(postData)
	return response.PostsResponse(postData)
}
