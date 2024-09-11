package request

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4,containsany=!@#$%"`
	Age      int    `json:"age" binding:"required,min=0,max=125,numeric"`
}
