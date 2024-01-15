package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(r *gin.Engine, db *gorm.DB) {
	SetUserRoutes(r, db)

	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})
}
