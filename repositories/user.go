package repositories

import (
	"waysbuck/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, ID int) (models.User, error)
	DeleteUser(user models.User, ID int) (models.User, error)
}

// type repository struct {
// 	db *gorm.DB
// }

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	//err := r.db.Raw("SELECT * FROM users").Scan(&users).Error
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	//err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error
	err := r.db.First(&user, ID).Error
	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	//err := r.db.Exec("INSERT INTO users(name,email,password,created_at,updated_at) VALUES (?,?,?,?,?)", user.Name, user.Email, user.Password, time.Now(), time.Now()).Error
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
	//rr := r.db.Raw("UPDATE users SET name=?, email=?, password=? WHERE id=?", user.Name, user.Email, user.Password, ID).Scan(&user).Error
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
	//err := r.db.Raw("DELETE FROM users WHERE id=?", ID).Scan(&user).Error
	err := r.db.Delete(&user).Error
	return user, err
}
