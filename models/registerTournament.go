package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type  Team struct {
	UserName		string					`json:"UserName" validate:"required"`
	GameUserName	string					`json:"GameUserName" validate:"required"`
	User_id			string					`json:"User_id" validate:"required"`
}

type RegisterTournament struct{
	ID						primitive.ObjectID		`bson:"_id" validate:"required"`
	Players					[]Player				`json:"Players" validate:"required"`
	Created_at				time.Time				`json:"Created_at" validate:"required"`
	Updated_at				time.Time				`json:"Updated_at" validate:"required"`
	TournamentId			string					`json:"TournamentId" validate:"required"`
	TournamentName			string					`json:"TournamentName" validate:"required"`
	TournamentIcon			string					`json:"TournamentIcon" validate:"required"`
	TournamentDate			string					`json:"TournamentDate" validate:"required"`
	RegisterTournamentId	string					`json:"RegisterTournamentId" validate:"required"`
	TeamName				string					`json:"TeamName" validate:"required"`
}