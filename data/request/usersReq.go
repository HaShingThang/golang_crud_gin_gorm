package request

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=2,max=100"`
}

type UpdateUserRequest struct {
	Id       int    `validate:"required"`
	Username string `json:"username" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=2,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=2,max=100"`
}
