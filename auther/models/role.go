package models

import (
	"auther/initializer"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}

// Create a role
func CreateRole(Role *Role) (err error) {
	err = initializer.DB.Create(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all roles
func GetRoles(Role *[]Role) (err error) {
	err = initializer.DB.Find(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get role by id
func GetRole(Role *Role, id int) (err error) {
	err = initializer.DB.Where("id = ?", id).First(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Update role
func UpdateRole(Role *Role) (err error) {
	initializer.DB.Updates(Role)
	return nil
}
