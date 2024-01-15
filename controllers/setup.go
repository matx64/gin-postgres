package controllers

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine) {
	SetUserRoutes(r)

	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})
}
