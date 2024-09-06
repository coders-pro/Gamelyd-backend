package queries

import (
	"context"
	"time"

	"github.com/Gameware/database"
	"github.com/Gameware/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var drawCollection *mongo.Collection = database.OpenCollection(database.Client, "draw")

func SaveMultiDraws(draws models.AllDraws) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var submitValue []interface{}
	for _, data := range draws {
		submitValue = append(submitValue, data)
	}

	_, insertErr := drawCollection.InsertMany(ctx, submitValue)
	return insertErr

}

func GetCurrentStageTeamsQuery(tournamentId string, stage int) (models.AllDraws, error) {

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

func GetSingleDrawQuery(id string) (models.Draw, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var draw models.Draw

	// Convert the string ID to primitive.ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return draw, err // Return if the conversion fails
	}

	// Perform the query
	err = drawCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&draw)
	return draw, err
}

func UpdateDrawQuery(id string, draw models.Draw) (*mongo.SingleResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"drawid": id}

	// Use the tournament struct to update the fields
	update := bson.M{
		"$set": draw,
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := drawCollection.FindOneAndUpdate(ctx, filter, update, &opt)

	// Decode the result into a Tournament struct
	var updatedDraw models.Draw
	if err := result.Decode(&updatedDraw); err != nil {
		return nil, err // Return error if decoding fails
	}

	return result, nil // Return the updated document and no error

}
