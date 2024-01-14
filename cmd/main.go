package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matx64/gin-postgres/cmd/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})
	r.GET("/users", controllers.ListUsersHandler)
	r.Run()
}
