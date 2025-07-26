package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost",
			"http://127.0.0.1",
			"https://api-go.ahmadsaubani.com",
			"https://ahmadsaubani.com",
			"https://www.ahmadsaubani.com",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// func CorsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		origin := c.GetHeader("Origin")
// 		if origin != "" {
// 			if strings.HasPrefix(origin, "http://localhost") ||
// 				strings.HasPrefix(origin, "http://127.0.0.1") ||
// 				strings.HasSuffix(origin, ".ahmadsaubani.com") {
// 				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
// 				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
// 				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 			}
// 		}

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusOK)
// 			return
// 		}

// 		c.Next()

// 	}
// }
