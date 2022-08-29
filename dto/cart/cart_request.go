package cartdto

type CreateCartRequest struct {
	ID        int      `json:"id" gorm:"primary_key:auto_increment"`
	ProductID int      `json:"product_id"`
	Products  []string `json:"product" gorm:"one2many:products"`
	Qty       int      `json:"qty" form:"qty"`
	// User      models.CartUserResponse      `json:"user"`
	//Topings   []models.TopingCartResponse  `json:"toping" gorm:"one2many:topings"`
}
