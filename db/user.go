package db

import (
	"github.com/google/uuid"
	"github.com/matx64/gin-postgres/models"
)

func CreateUser(user models.User) (uuid.UUID, error) {
	result := DB.Create(user)

	return user.ID, result.Error
}

func ListUsers() ([]models.User, error) {
	users := []models.User{}

	result := DB.Find(&users)

	return users, result.Error
}

func GetUser(uuid uuid.UUID) (models.User, error) {
	user := models.User{}

	result := DB.First(&user, uuid)

	return user, result.Error
}

func UpdateUser(uuid uuid.UUID, input models.UserCreate) error {
	user, err := GetUser(uuid)

	if err != nil {
		return err
	}

	updatedUser := models.User{Name: input.Name, Email: input.Email, Password: input.Password}

	result := DB.Model(&user).Updates(&updatedUser)

	return result.Error
}

func DeleteUser(uuid uuid.UUID) error {
	_, err := GetUser(uuid)

	if err != nil {
		return err
	}

	result := DB.Delete(&models.User{}, uuid)

	return result.Error
}
