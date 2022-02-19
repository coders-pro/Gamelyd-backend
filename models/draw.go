package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type  Player struct {
	UserName		string					`json:"UserName" validate:"required"`
	GameUserName	string					`json:"GameUserName" validate:"required"`
	User_id			string					`json:"User_id" validate:"required"`
}

type  Teams struct {
	TeamName		string					`json:"TeamName" validate:"required"`
	Players			[]Player				`json:"Players" validate:"required"`
}

type Draw struct{
	ID						primitive.ObjectID		`bson:"_id" validate:"required"`
	Team1					Teams					`json:"Team1" validate:"required"`
	Team2					Teams					`json:"Team2" validate:"required"`
	Created_at				time.Time				`json:"Created_at" validate:"required"`
	Updated_at				time.Time				`json:"Updated_at" validate:"required"`
	TournamentId			string					`json:"TournamentId" validate:"required"`
	DrawId					string					`json:"DrawId" validate:"required"`
	Stage					int						`json:"Stage" validate:"required"`
	Winner					string					`json:"Winner" validate:"eq=Team1|eq=Team2"`
	Time					string					`json:"Time"`
	Date					string					`json:"Date"`
	Team1Score				int						`json:"Team1Score"`
	Team2Score				int						`json:"Team2Score"`
	Link					string					`json:"Link"`

}


