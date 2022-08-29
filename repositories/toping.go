package repositories

import (
	"waysbuck/models"

	"gorm.io/gorm"
)

type TopingRepository interface {
	FindTopings() ([]models.Toping, error)
	GetToping(ID int) (models.Toping, error)
	CreateToping(toping models.Toping) (models.Toping, error)
	UpdateToping(toping models.Toping, ID int) (models.Toping, error)
	DeleteToping(toping models.Toping, ID int) (models.Toping, error)
}

func RepositoryToping(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTopings() ([]models.Toping, error) {
	var topings []models.Toping
	//err := r.db.Raw("SELECT * FROM users").Scan(&users).Error
	err := r.db.Find(&topings).Error
	return topings, err
}

func (r *repository) GetToping(ID int) (models.Toping, error) {
	var toping models.Toping
	//err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error
	err := r.db.First(&toping, ID).Error
	return toping, err
}

func (r *repository) CreateToping(toping models.Toping) (models.Toping, error) {
	//err := r.db.Exec("INSERT INTO users(name,email,password,created_at,updated_at) VALUES (?,?,?,?,?)", user.Name, user.Email, user.Password, time.Now(), time.Now()).Error
	err := r.db.Create(&toping).Error
	return toping, err
}

func (r *repository) UpdateToping(toping models.Toping, ID int) (models.Toping, error) {
	err := r.db.Save(&toping).Error
	return toping, err
}

func (r *repository) DeleteToping(toping models.Toping, ID int) (models.Toping, error) {
	err := r.db.Delete(&toping).Error
	return toping, err
}
