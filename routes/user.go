package routes

import (
	controller "github.com/Gameware/controllers"
	"github.com/Gameware/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users/checkUserName/:name", controller.CheckUserName())

	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/editUser/:user_id", controller.UpdateUser())
	incomingRoutes.GET("/users/delete/:user_id", controller.DeleteUser())
	incomingRoutes.GET("/users/tournaments/:id", controller.ListUserTournament())

}