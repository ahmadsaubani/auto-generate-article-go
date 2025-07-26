package routes

import (
	"net/http"
	"news-go/src/configs/database"
	"news-go/src/controllers/api/v1/article_controllers"
	"news-go/src/middlewares"
	"news-go/src/repositories/article_repositories"
	"news-go/src/services/article_services"

	"github.com/gin-gonic/gin"
)

func API(db *database.DBConnection, ginEngine *gin.Engine) *gin.Engine {

	articleRepository := article_repositories.NewArticleRepository()
	articleService := article_services.NewArticleService(articleRepository)

	ginEngine.Use(middlewares.CorsMiddleware())
	v1 := ginEngine.Group("/api/v1")
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		v1.GET("/article/pulls", article_controllers.FetchArticles(articleService))
		v1.GET("/articles", article_controllers.GetAllArticles(articleService))

	}

	return ginEngine
}
