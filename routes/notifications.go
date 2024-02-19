package routes

import (
	controller "github.com/Gameware/controllers"
	"github.com/Gameware/middleware"
	"github.com/gin-gonic/gin"
)

func NotificationsRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/notifications", controller.CreateNotification())
	incomingRoutes.GET("/notifications/:userID", controller.GetNotifications())
	incomingRoutes.PUT("/notifications/:id", controller.MarkNotificationToIsRead())
	incomingRoutes.PUT("/notifications/markall/:userID", controller.MarkAllNotificationsAsRead())
	incomingRoutes.DELETE("/notifications/:id", controller.DeleteNotification())
}
