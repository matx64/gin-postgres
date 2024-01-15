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

func SetUserRoutes(r *gin.Engine, dbInstance *gorm.DB) {
	g := r.Group("/users")
	g.POST("", createUserHandler(dbInstance))
	g.GET("", listUsersHandler(dbInstance))
	g.GET("/:id", getUserHandler(dbInstance))
	g.PATCH("/:id", updateUserHandler(dbInstance))
	g.DELETE("/:id", deleteUserHandler(dbInstance))
}

func createUserHandler(dbInstance *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		input := new(models.UserCreate)

		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := models.NewUser(*input)

		id, err := db.CreateUser(dbInstance, user)

		if err != nil {
			handleDbError(c, err)
			return
		}

		c.JSON(http.StatusCreated, id)
	}
}

func listUsersHandler(dbInstance *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := db.ListUsers(dbInstance)

		if err != nil {
			handleDbError(c, err)
			return
		}

		c.JSON(http.StatusOK, users)
	}

}

func getUserHandler(dbInstance *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, err := uuid.Parse(c.Param("id"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := db.GetUser(dbInstance, uuid)

		if err != nil {
			handleDbError(c, err)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func updateUserHandler(dbInstance *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		if err := db.UpdateUser(dbInstance, uuid, *input); err != nil {
			handleDbError(c, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func deleteUserHandler(dbInstance *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, err := uuid.Parse(c.Param("id"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.DeleteUser(dbInstance, uuid); err != nil {
			handleDbError(c, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func handleDbError(c *gin.Context, err error) {
	status := http.StatusInternalServerError

	if errors.Is(err, gorm.ErrRecordNotFound) {
		status = http.StatusNotFound
	}

	c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
}
