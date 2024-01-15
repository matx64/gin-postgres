package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matx64/gin-postgres/controllers"
	"github.com/matx64/gin-postgres/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDb()

	r := gin.Default()
	controllers.SetRoutes(r)
	r.Run(":1337")
}
