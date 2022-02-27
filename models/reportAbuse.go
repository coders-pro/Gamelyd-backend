package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type ReportAbuse struct{
	ID						primitive.ObjectID		`bson:"_id" validate:"required"`
	Created_at				time.Time				`json:"Created_at" validate:"required"`
	Updated_at				time.Time				`json:"Updated_at" validate:"required"`
	ReportAbuseId			string					`json:"ContactId" validate:"required"`
	UserID					string					`json:"UserID" validate:"required"`
	Email					string					`json:"Email" validate:"required"`
	Name					string					`json:"Name" validate:"required"`
	Message					string					`json:"Message" validate:"required"`
	For						string					`json:"For" validate:"required"`
	Achived					bool					`json:"Achived"`
	IsDeleted				bool					`json:"IsDeleted"`
	IsCompleted				bool					`json:"IsCompleted"`
}


