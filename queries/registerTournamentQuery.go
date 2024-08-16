package queries

import (
	"context"
	"time"

	"github.com/Gameware/database"
	"github.com/Gameware/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var registerTournamentCollection *mongo.Collection = database.OpenCollection(database.Client, "registerTournament")

func GetRegisteredTeamsQuery(id string) (models.AllRegTeams, error)  {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var teams models.AllRegTeams

	cursor, err := registerTournamentCollection.Find(ctx, bson.M{"tournamentid": id})
	if err != nil {
		return nil, err // Handle the error appropriately
	}
	defer cursor.Close(ctx) // Ensure the cursor is closed after processing

	if err := cursor.All(ctx, &teams); err != nil {
		return nil, err
	}

	return teams, err

}