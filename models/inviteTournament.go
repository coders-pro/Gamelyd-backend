package models

import (
	"time"
)

type InviteTournament struct {
	Tournament_id string    `bson:"tournament_id" validate:"required"`
	User_id       string    `json:"user_id" validate:"required"`
	Created_at    time.Time `json:"CreatedAt" validate:"required"`
	Updated_at    time.Time `json:"UpdatedAt" validate:"required"`
}
