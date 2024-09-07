package queries

import (
	"context"
	"time"

	"github.com/Gameware/database"
	"github.com/Gameware/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tournamentCollection *mongo.Collection = database.OpenCollection(database.Client, "tournament")

func GetSingleTournamentQuery(id string) (models.Tournament, error)  {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var tournament models.Tournament
	err := tournamentCollection.FindOne(ctx, bson.M{"tournamentid": id}).Decode(&tournament)
	return tournament, err

}

func StartTournamentQuery(id string) *mongo.SingleResult  {
	
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"tournamentid": id }
	update := bson.M{
		"$set": bson.M{"start": true},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert, 
	}

	result := tournamentCollection.FindOneAndUpdate(ctx, filter, update, &opt)

	return result

}

func UpdateTournamentQuery(id string, tournament models.Tournament) (*mongo.SingleResult, error) {
    var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
    defer cancel()

    filter := bson.M{"tournamentid": id}

    // Use the tournament struct to update the fields
    update := bson.M{
        "$set": tournament,
    }

    upsert := true
    after := options.After
    opt := options.FindOneAndUpdateOptions{
        ReturnDocument: &after,
        Upsert:         &upsert,
    }

    result := tournamentCollection.FindOneAndUpdate(ctx, filter, update, &opt)

	   // Decode the result into a Tournament struct
	   var updatedTournament models.Tournament
	   if err := result.Decode(&updatedTournament); err != nil {
		   return nil, err // Return error if decoding fails
	   }
   
	   return result, nil // Return the updated document and no error

}
