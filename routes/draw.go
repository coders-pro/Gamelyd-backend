package routes

import (
	controller "github.com/Gameware/controllers"
	"github.com/Gameware/middleware"
	"github.com/gin-gonic/gin"
)

func DrawRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/draws/save", controller.Draw())
	incomingRoutes.GET("/draws/:tornamentId", controller.GetDrawByTornamentID())
	incomingRoutes.POST("/draws/addWinner/:drawId", controller.AddWinner())
	incomingRoutes.POST("/draws/addTime/:drawId", controller.AddTime())
	incomingRoutes.POST("/draws/addScore/:drawId", controller.AddScore())
}