package migration

import (
	"auther/initializer"
	"auther/models"
	"os"
)

func MigrateDB() {
	initializer.DB.AutoMigrate(&models.Role{})
	initializer.DB.AutoMigrate(&models.User{})
	seedData()
}

// load seed data into the database
func seedData() {
	var roles = []models.Role{{Name: "admin", Description: "Administrator role"}, {Name: "operator", Description: "Authenticated operator role"}, {Name: "user", Description: "Authenticated user role"}}

	// creat admin user	
	user := models.User{Username: os.Getenv("ADMIN_USERNAME"),
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		RoleID:   1}
	
	initializer.DB.Save(&roles)
	_, err := user.Save()
	if err != nil {
		return
	}
}
