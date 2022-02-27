package routes

import (
	controller "github.com/Gameware/controllers"
	"github.com/Gameware/middleware"
	"github.com/gin-gonic/gin"
)

func ContactUsRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/contactUs/save", controller.SaveContactUs())
	incomingRoutes.GET("/contactUs/getAll", controller.GetAllContactUs())
	
}