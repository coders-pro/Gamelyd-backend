package routes

import (
	controller "github.com/Gameware/controllers"
	"github.com/Gameware/middleware"
	"github.com/gin-gonic/gin"
)

func DrawRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/draw/save", controller.Draw())
	incomingRoutes.GET("/draws/:tornamentId", controller.GetDrawByTornamentID())
}