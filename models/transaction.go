package models

import "time"

type Transaction struct {
	ID      int                  `json:"id" gorm:"primary_key:auto_increment"`
	BuyerID int                  `json:"buyer_id"`
	Buyer   UsersProfileResponse `json:"buyer"`
	Price   int                  `json:"price"`
	Product Product               `json:"product_id"`
	Total     int       `json:"total"`
	Status    string    `json:"status"  gorm:"type:varchar(25)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type TransactionResponse struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	BuyerID   int                  `json:"buyer_id"`
	Buyer     UsersProfileResponse `json:"buyer"`
	Price     int                  `json:"price"`
	Status    string               `json:"status"  gorm:"type:varchar(25)"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
