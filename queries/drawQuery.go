package queries

import (
	"context"
	"time"

	"github.com/Gameware/database"
	"github.com/Gameware/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var drawCollection *mongo.Collection = database.OpenCollection(database.Client, "draw")

func SaveMultiDraws(draws models.AllDraws) error  {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var submitValue []interface{}
	for _, data := range draws {
		submitValue = append(submitValue, data)
	}

	_, insertErr := drawCollection.InsertMany(ctx, submitValue)
	return insertErr

}

func GetCurrentStageTeamsQuery(tournamentId string, stage int) (models.AllDraws, error)  {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var draws models.AllDraws

	cursor, err := drawCollection.Find(ctx, bson.M{"stage": stage, "tournamentid": tournamentId})
	if err != nil {
		return nil, err // Handle the error appropriately
	}
	defer cursor.Close(ctx) // Ensure the cursor is closed after processing

	if err := cursor.All(ctx, &draws); err != nil {
		return nil, err
	}

	return draws, err

}