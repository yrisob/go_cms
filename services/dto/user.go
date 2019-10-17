package dto

//UserDTO data transfer object for getting by form
type UserDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Role     int    `json:"role"`
	Phone    string `json:"phone" validate:"phone"`
	Password string `json:"password" validate:"password"`
}
