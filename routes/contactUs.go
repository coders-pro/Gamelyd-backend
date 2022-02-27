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
	incomingRoutes.GET("/contactUs/delete/:id", controller.DeleteContact())
	incomingRoutes.GET("/contactUs/complete/:id", controller.CompleteContact())
	incomingRoutes.GET("/contactUs/achive/:id", controller.AchivedContact())
	incomingRoutes.GET("/contactUs/active", controller.GetActiveContact())
	incomingRoutes.GET("/contactUs/deleted", controller.GetIsDeletedContact())
	incomingRoutes.GET("/contactUs/completed", controller.GetCompleteContact())
	incomingRoutes.GET("/contactUs/achived", controller.GetIsAchivedContact())
}