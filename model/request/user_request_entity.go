package request

type UserRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=50"`
	Age    int    `json:"age" validate:"required,numeric"`
	Email  string `json:"email" validate:"required,email"`
	Gender string `json:"gender" validate:"required,alpha"`
}
