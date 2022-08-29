package models

import "time"

type Cart struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	ProductID []int     `json:"product_id" form:"product_id" gorm:"-"`
	Product   []Product `json:"product" gorm:"many2many:product"`
	Qty       int       `json:"qty" form:"qty"`
	// UserID    int                   `json:"user_id" form:"user_id" gorm:"-"`
	// User      UsersProfileResponse  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TopingID  []int                `json:"toping_id" form:"toping_id" gorm:"-"`
	Toping    []TopingCartResponse `json:"toping" gorm:"many2many:topings"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type CartResponse struct {
	ID        int                  `json:"id"`
	ProductID []int                `json:"product_id" gorm:"-"`
	Product   []Product            `json:"product" form:"product" gorm:"many2many:product"`
	Qty       int                  `json:"qty"`
	UserID    int                  `json:"-"`
	TopingID  []int                `json:"toping_id" gorm:"-"`
	Toping    []TopingCartResponse `json:"toping" gorm:"many2many:topings"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type CartUserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (CartResponse) TableName() string {
	return "carts"
}

func (CartUserResponse) TableName() string {
	return "carts"
}
