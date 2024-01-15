package db

import (
	"github.com/google/uuid"
	"github.com/matx64/gin-postgres/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user models.User) (uuid.UUID, error) {
	result := db.Create(user)

	return user.ID, result.Error
}

func ListUsers(db *gorm.DB) ([]models.User, error) {
	users := []models.User{}

	result := db.Find(&users)

	return users, result.Error
}

func GetUser(db *gorm.DB, uuid uuid.UUID) (models.User, error) {
	user := models.User{}

	result := db.First(&user, uuid)

	return user, result.Error
}

func UpdateUser(db *gorm.DB, uuid uuid.UUID, input models.UserCreate) error {
	user, err := GetUser(db, uuid)

	if err != nil {
		return err
	}

	updatedUser := models.User{Name: input.Name, Email: input.Email, Password: input.Password}

	result := db.Model(&user).Updates(&updatedUser)

	return result.Error
}

func DeleteUser(db *gorm.DB, uuid uuid.UUID) error {
	_, err := GetUser(db, uuid)

	if err != nil {
		return err
	}

	result := db.Delete(&models.User{}, uuid)

	return result.Error
}
