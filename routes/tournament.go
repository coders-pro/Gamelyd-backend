package routes

import (
	controller "github.com/Gameware/controllers"
	"github.com/Gameware/middleware"
	"github.com/gin-gonic/gin"
)

func TournamentRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/tournament/save", controller.SaveTournament())
	incomingRoutes.GET("/tournaments/:search/:page", controller.GetTournaments())
	incomingRoutes.GET("/tournament/:id", controller.GetTournament())
	incomingRoutes.GET("tournament/mode/:paymentType/limit", controller.GetTournamentByType())
	incomingRoutes.GET("/tournament/delete/:id", controller.DeleteTournament())
	incomingRoutes.GET("/tournament/suspend/:id", controller.SuspendTournament())
	incomingRoutes.POST("/tournament/update/:id", controller.UpdateTournament())
	incomingRoutes.GET("/tournament/participants/:id", controller.ListPartTournament())
	incomingRoutes.POST("/tournament/register/:tournamentId", controller.RegisterTournament())
	incomingRoutes.GET("/tournament/mode/:paymentMode/:page", controller.GetTournamentsByMode())
	incomingRoutes.GET("/tournament/mode/:paymentMode/limit", controller.GetTournamentsByModeLimit())
	incomingRoutes.GET("/tournament/userRegisteredTournaments/:id/:page", controller.UserTournaments())
	incomingRoutes.GET("/tournament/userRegisteredTournaments/:id/limit", controller.UserTournamentsLimit())
	incomingRoutes.GET("/tournament/removeFromTournament/:userId/:tournamentId", controller.RemoveUser())

}