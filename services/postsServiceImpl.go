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
func (ps *PostsServiceImpl) Create(posts request.CreatePostRequest) response.PostResponse {
	err := ps.Validate.Struct(posts)
	helpers.ErrorHandler(err)
	Posts := models.Post{
		Title:       posts.Title,
		Description: posts.Description,
		UserId:      posts.UserId,
	}
	ps.PostsInterface.Save(Posts)
	return response.PostResponse(Posts)
}

// Delete implements PostsService.
func (ps *PostsServiceImpl) Delete(postsId int) {
	ps.PostsInterface.Delete(postsId)
}

// FindAll implements PostsService.
func (ps *PostsServiceImpl) FindAll() []response.PostResponse {
	result := ps.PostsInterface.FindAll()

	var posts []response.PostResponse
	for _, value := range result {
		post := response.PostResponse{
			Id:          value.Id,
			Title:       value.Title,
			Description: value.Description,
			UserId:      value.UserId,
		}
		posts = append(posts, post)
	}

	return posts
}

// FindById implements PostsService.
func (ps *PostsServiceImpl) FindById(postsId int) response.PostResponse {
	postData, err := ps.PostsInterface.FindById(postsId)
	helpers.ErrorHandler(err)

	postRes := response.PostResponse{
		Id:          postData.Id,
		Title:       postData.Title,
		Description: postData.Description,
		UserId:      postData.UserId,
	}
	return postRes
}

// Update implements PostsService.
func (ps *PostsServiceImpl) Update(posts request.UpdatePostRequest) response.PostResponse {
	postData, err := ps.PostsInterface.FindById(posts.Id)
	helpers.ErrorHandler(err)
	postData.Title = posts.Title
	postData.Description = posts.Description
	ps.PostsInterface.Update(postData)
	return response.PostResponse(postData)
}
