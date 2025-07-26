package main

import (
	"news-go/src/configs/database"
	"news-go/src/routes"
	"news-go/src/seeders"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()
	ginEngine := gin.Default()

	ginEngine.Use(cors.New(cors.Config{
		// AllowOrigins: []string{
		// 	"http://localhost",
		// 	"http://127.0.0.1",
		// 	"https://api-go.ahmadsaubani.com",
		// 	"https://ahmadsaubani.com",
		// 	"https://www.ahmadsaubani.com",
		// },
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// hidupkan jika development
	// if err := godotenv.Load(); err != nil {
	// 	panic("Error loading .env file: " + err.Error())
	// }

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
