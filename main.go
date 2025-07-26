package main

import (
	"news-go/src/configs/database"
	"news-go/src/routes"
	"news-go/src/seeders"

	"github.com/gin-gonic/gin"
//	"github.com/joho/godotenv"
)

func main() {
	gin.DisableConsoleColor()
	ginEngine := gin.Default()

//	if err := godotenv.Load(); err != nil {
//		panic("Error loading .env file: " + err.Error())
//	}

	db := database.ConnectDatabase()

	// Run migrations
	database.RunMigrations(db.Gorm)

	seeders.Run(db)
	r := routes.API(db, ginEngine)
	if err := r.Run(":3000"); err != nil {
		panic("Error starting server: " + err.Error())
	}

	println("Server is running...")
}
