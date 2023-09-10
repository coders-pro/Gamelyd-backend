package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"user_id"`
	Message   string             `bson:"message"`
	Timestamp time.Time          `bson:"timestamp"`
	IsRead    bool               `bson:"is_read"`
}
