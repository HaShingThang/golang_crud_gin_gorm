package request

type CreatePostRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=80"`
	Description string `json:"description" validate:"required,min=1,max=1500"`
	UserId      int    `json:"userId" validate:"required"`
}

type UpdatePostRequest struct {
	Id          int    `validate:"required"`
	Title       string `json:"title" validate:"required,min=1,max=80"`
	Description string `json:"description" validate:"required,min=1,max=1500"`
	UserId      int    `json:"userId" validate:"required"`
}
