package database

import (
	"fmt"
	"waysbuck/models"
	mysql "waysbuck/pkg"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Toping{},
		&models.Transaction{},
		&models.Profile{},
		&models.Cart{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
