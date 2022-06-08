package routes

import (
	"github.com/FredsZDev/WebApi_GO_B2/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) *gin.Engine {
	main := router.Group("WebApi/b2")
	{
		games := main.Group("games")
		{
			games.GET("/:id", controllers.ShowGames)
			games.GET("/:", controllers.ShowGames)
			games.POST("/", controllers.CreateGame)
			games.PUT("/", controllers.UpdateGames)
			games.DELETE("/:ID", controllers.DeleteGames)
		}

	}
	return router
}
