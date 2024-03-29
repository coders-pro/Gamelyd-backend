package routes

import (
	controller "github.com/Gameware/controllers"
	"github.com/Gameware/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users/checkUserName/:name", controller.CheckUserName())
	incomingRoutes.GET("/test", controller.Test())

	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/changePassword/:user_id", controller.ChangePassword())
	incomingRoutes.POST("/users/editUser/:user_id", controller.UpdateUser())
	incomingRoutes.POST("/users/reset", controller.ResetPassword())
	incomingRoutes.GET("/users/delete/:user_id", controller.DeleteUser())
	incomingRoutes.GET("/users/tournaments/:id/:page", controller.ListUserTournament())
	incomingRoutes.GET("/users/tournaments/:id/limit", controller.ListUserTournamentLimit())
	incomingRoutes.GET("/users/search/:search/:page", controller.SearchUsers())

}
