package models

import "time"

type Toping struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc      string    `json:"desc" gorm:"type:text" form:"desc"`
	Price     int       `json:"price" form:"price" gorm:"type: int"`
	Image     string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty       int       `json:"qty" form:"qty"`
	Cart      []Cart    `json:"cart" gorm:"many2many:carts"`
	Product   []Product `json:"product" gorm:"many2many:product"`
	UserID    int       `json:"user_id" form:"user_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type TopingResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	Qty    int    `json:"qty"`
	UserID int    `json:"-"`
}

type TopingCartResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Qty   int    `json:"qty"`
}

func (TopingResponse) TableName() string {
	return "products"
}
