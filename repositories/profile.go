package repositories

import (
	"waysbuck/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindProfiles() ([]models.Profile, error)
	GetProfile(ID int) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProfiles() ([]models.Profile, error) {
	var profiles []models.Profile
	//err := r.db.Raw("SELECT * FROM users").Scan(&users).Error
	err := r.db.Find(&profiles).Error
	return profiles, err
}

func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var profile models.Profile
	//err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error
	err := r.db.First(&profile, ID).Error
	return profile, err
}
