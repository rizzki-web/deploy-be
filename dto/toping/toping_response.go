package topingsdto

type TopingResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" form:"name" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
}
