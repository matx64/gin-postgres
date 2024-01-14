package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/matx64/gin-postgres/cmd/models"
)

func ListUsersHandler(c *gin.Context) {
	users := [1]models.User{models.NewUser(1, "Matheus")}

	c.JSON(200, users)
}
