package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		// AllowOrigins: []string{
		// 	"http://localhost",
		// 	"http://127.0.0.1",
		// 	"https://api-go.ahmadsaubani.com",
		// 	"https://ahmadsaubani.com",
		// 	"https://www.ahmadsaubani.com",
		// },
		AllowOrigins:     []string{"*"}, // Allow all origins for development purposes
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
