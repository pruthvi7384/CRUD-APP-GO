package repositories

import (
	"crudApp/config"
	"crudApp/models"
)

// Create User Func
func CreateUser(user models.User) (error, models.User) {
	err := config.DB.Create(&user).Error
	return err, user
}

// Find User By Email Id Func
func FindUser(email string) (error, models.User) {
	var user = models.User{}
	err := config.DB.Where("email = ?", email).First(&user).Error
	return err, user
}

// Remove User Record
func RemoveUser(id uint) error {
	err := config.DB.Delete(&models.User{}, id).Error
	return err
}

// Find User By Id Func
func FindUserById(id uint) (error, models.User) {
	user := models.User{}
	err := config.DB.Where("id = ?", id).First(&user).Error
	return err, user
}

// Update User Func
func UpdateUser(user models.User) (error, models.User) {
	err := config.DB.Save(&user).Error
	return err, user
}
