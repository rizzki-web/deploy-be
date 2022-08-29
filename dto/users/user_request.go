package usersdto

type CreateUserRequest struct {
	Name     string   `json:"name" form:"name" validate:"required"`
	Email    string   `json:"email" form:"email" validate:"required"`
	Password string   `json:"password" form:"password" validate:"required"`
	Product  []string `json:"product" gorm:"one2many:product"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}