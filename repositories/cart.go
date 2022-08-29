package repositories

import (
	"waysbuck/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCart() ([]models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	GetProductCart(ID int) (models.Product, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart, ID int) (models.Cart, error)
	CleaningCart(cart models.Cart) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCart() ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Product").Preload("Toping").Preload("User").Find(&cart).Error
	return cart, err
}

func (r *repository) GetProductCart(ID int) (models.Product, error) {
	var productCart models.Product
	err := r.db.First(&productCart, ID).Error
	return productCart, err
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.First(&cart, ID).Error
	return cart, err
}

// *NOTE: buat array untuk menampung modelsProduct, modelsCart, errorProduct, errorCart
// agar nantinnya dapat direturn
func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error
	return cart, err

}

func (r *repository) DeleteCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Delete(&cart).Error
	return cart, err
}

func (r *repository) CleaningCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error
	return cart, err
}
