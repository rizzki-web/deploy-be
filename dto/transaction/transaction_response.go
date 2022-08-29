package transactionsdto

//import category "waysbuck/models"

type TransactionResponse struct {
	Price   int `json:"price" form:"price" gorm:"type: int"`
	BuyerID int `json:"buyer_id" gorm:"type: int"`
}
