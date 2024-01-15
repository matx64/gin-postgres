package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/matx64/gin-postgres/db"
	"github.com/matx64/gin-postgres/models"
	"gorm.io/gorm"
)

func SetUserRoutes(r *gin.Engine) {
	g := r.Group("/users")
	g.POST("", CreateUserHandler)
	g.GET("", ListUsersHandler)
	g.GET("/:id", GetUserHandler)
	g.PATCH("/:id", UpdateUserHandler)
	g.DELETE("/:id", DeleteUserHandler)
}

func CreateUserHandler(c *gin.Context) {
	input := new(models.UserCreate)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.NewUser(*input)

	id, err := db.CreateUser(user)

	if err != nil {
		handleDbError(c, err)
		return
	}

	c.JSON(http.StatusCreated, id)
}

func ListUsersHandler(c *gin.Context) {
	users, err := db.ListUsers()

	if err != nil {
		handleDbError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserHandler(c *gin.Context) {
	uuid, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := db.GetUser(uuid)

	if err != nil {
		handleDbError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserHandler(c *gin.Context) {
	uuid, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := new(models.UserCreate)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.UpdateUser(uuid, *input)

	if err != nil {
		handleDbError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func DeleteUserHandler(c *gin.Context) {
	uuid, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.DeleteUser(uuid)

	if err != nil {
		handleDbError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func handleDbError(c *gin.Context, err error) {
	status := http.StatusInternalServerError

	if errors.Is(err, gorm.ErrRecordNotFound) {
		status = http.StatusNotFound
	}

	c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
}
