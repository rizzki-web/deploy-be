package topingsdto

//import category "waysbuck/models"

type CreateTopingRequest struct {
	Name  string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc  string `json:"desc" gorm:"type:text" form:"desc"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty   int    `json:"qty" form:"qty" gorm:"type: int"`
}

type UpdateTopingRequest struct {
	Name  string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc  string `json:"desc" gorm:"type:text" form:"desc"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty   int    `json:"qty" form:"qty" gorm:"type: int"`
}
