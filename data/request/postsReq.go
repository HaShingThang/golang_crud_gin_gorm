package request

type CreatePostsRequest struct {
	Title       string `validate:"required,min=1,max=80" json:"title"`
	Description string `validate:"required,min=1,max=1500" json:"description"`
}

type UpdatePostsRequest struct {
	Id          int    `validate:"required"`
	Title       string `validate:"required,min=1,max=80" json:"title"`
	Description string `validate:"required,min=1,max=1500" json:"description"`
}

