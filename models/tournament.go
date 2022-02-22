package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tournament struct{
	ID							primitive.ObjectID		`bson:"_id" validate:"required"`
	Name						*string					`json:"Name" validate:"required,min=2,max=100"`
	GameName					*string					`json:"GameName" validate:"required,min=2,max=100"`
	Icon						*string					`json:"icon" validate:"required,min=1,max=100"`
	TournamentType				*string					`json:"TournamentType" validate:"required,eq=PUBLIC|eq=PRIVATE"`
	Payment						*string					`json:"Payment" validate:"required,eq=FREE|eq=PAID|eq=SPONSORED"`
	TournamentMode				*string					`json:"TournamentMode" validate:"required,eq=MULTIPLAYER|eq=BATTLEROYALE"`
	TournamentSize				*int			    	`json:"TournamentSize" validate:"required"`
	Team						*string					`json:"Team" validate:"required,eq=SINGLE|eq=DUO|eq=SQUAD"`
	Shuffle						*string					`json:"Shuffle" validate:"required,eq=MANUAL|eq=AUTOMATIC"`
	Created_at					time.Time				`json:"Created_at" validate:"required"`
	Updated_at					time.Time				`json:"Updated_at"`
	TournamentId				string					`json:"TournamentId" validate:"required"`
	User_id						string					`json:"User_id" validate:"required"`
	Active						bool					`json:"Active"`
	IsDeleted					bool					`json:"IsDeleted"`
	IsSuspended					bool					`json:"IsSuspended"`
	Start						bool					`json:"Start"`
	Link						string					`json:"Link"`
	Date						string					`json:"Date"`
	IsPaid						bool					`json:"IsPaid"`
	RefNumber					string					`json:"RefNumber"`
	PaymentChannel				string					`json:"PaymentChannel"`
	Amount						int						`json:"Amount"`
	RegistrationAmount			int						`json:"RegistrationAmount"`
}