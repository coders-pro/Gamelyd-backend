package routes

import (
	controller "github.com/Gameware/controllers"
	"github.com/Gameware/middleware"
	"github.com/gin-gonic/gin"
)

func ReportAbuseRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/reportAbuse/save", controller.SaveReportAbuse())
	incomingRoutes.GET("/reportAbuse/getAll", controller.GetAllReportAbuse())
	incomingRoutes.GET("/reportAbuse/delete/:id", controller.DeleteReportAbuse())
	incomingRoutes.GET("/reportAbuse/complete/:id", controller.CompleteReportAbuse())
	incomingRoutes.GET("/reportAbuse/achive/:id", controller.AchivedReportAbuse())
	incomingRoutes.GET("/reportAbuse/active", controller.GetActiveReportAbuse())
	incomingRoutes.GET("/reportAbuse/deleted", controller.GetIsDeletedReportAbuse())
	incomingRoutes.GET("/reportAbuse/completed", controller.GetCompleteReportAbuse())
	incomingRoutes.GET("/reportAbuse/achived", controller.GetIsAchivedReportAbuse())
}