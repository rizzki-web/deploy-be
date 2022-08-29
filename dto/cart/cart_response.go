package cartdto

import "waysbuck/models"

type CartResponse struct {
	ID        int                          `json:"id" gorm:"primary_key:auto_increment"`
	Product   []models.ProductCartResponse `json:"product" gorm:"one2many:product"`
	ProductID []int                        `json:"-" form:"product_id" gorm:"-"`
	Qty       int                          `json:"qty" form:"qty"`
	UserID    int                          `json:"user_id" form:"user_id"`
	//Toping    []string                     `json:"toping" gorm:"many2many:topings"`
}
