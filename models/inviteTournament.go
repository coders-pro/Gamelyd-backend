package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InviteTournament struct {
	Tournament_id primitive.ObjectID `bson:"tournamrnt_id" validate:"required"`
	User_id       primitive.ObjectID `json:"user_id" validate:"required"`
	Created_at    time.Time          `json:"CreatedAt" validate:"required"`
	Updated_at    time.Time          `json:"UpdatedAt" validate:"required"`
}
