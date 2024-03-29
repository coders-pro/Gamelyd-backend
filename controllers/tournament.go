package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"

	"net/http"

	"time"

	"github.com/Gameware/database"
	"github.com/Gameware/templates"

	helper "github.com/Gameware/helpers"
	"github.com/Gameware/models"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tournamentCollection *mongo.Collection = database.OpenCollection(database.Client, "tournament")
var registerTournamentCollection *mongo.Collection = database.OpenCollection(database.Client, "registerTournament")
var inviteTournamentCollection *mongo.Collection = database.OpenCollection(database.Client, "inviteTournament")

func SaveTournament() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var tournament models.Tournament

		if err := c.BindJSON(&tournament); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		tournament.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		tournament.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		tournament.ID = primitive.NewObjectID()
		tournament.TournamentId = tournament.ID.Hex()
		tournament.User_id = c.GetString("uid")
		tournament.Active = false
		tournament.IsDeleted = false
		tournament.IsSuspended = false
		tournament.Start = false
		tournament.IsPaid = false

		validationErr := validate.Struct(tournament)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message": validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		resultInsertionNumber, insertErr := tournamentCollection.InsertOne(ctx, tournament)
		if insertErr != nil {
			msg := "item was not created"
			c.JSON(http.StatusOK, gin.H{"message": msg, "hasError": true})
			defer cancel()
			return
		}
		defer cancel()
		go helper.SendEmail(c.GetString("email"), templates.CreateTournament(c.GetString("first_name")+" "+c.GetString("last_name")), "New Tournament")

		c.JSON(http.StatusOK, gin.H{"message": "Tournament created successfully", "data": tournament, "hasError": false, "insertId": resultInsertionNumber})
	}
}

// userID 61e99c5efa54a7d01ff272ce
// tournamentID 61e6c5b175740deeec73b156
func RegisterTournament() gin.HandlerFunc {
	return func(c *gin.Context) {
		tournamentId := c.Param("tournamentId")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var registerTournament models.RegisterTournament

		if err := c.BindJSON(&registerTournament); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}
		registerTournament.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		registerTournament.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		registerTournament.ID = primitive.NewObjectID()
		registerTournament.RegisterTournamentId = registerTournament.ID.Hex()
		registerTournament.TournamentId = tournamentId
		registerTournament.TournamentId = tournamentId

		count, err := registerTournamentCollection.CountDocuments(ctx, bson.M{"teamname": registerTournament.TeamName, "tournamentid": registerTournament.TournamentId})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusOK, gin.H{"message": "error occured while checking for the team name", "hasError": true})
			return
		}

		if count > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "Team name is not available", "hasError": true})
			return
		}

		var tournament models.Tournament
		err = tournamentCollection.FindOne(ctx, bson.M{"tournamentid": tournamentId}).Decode(&tournament)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		counted, err := registerTournamentCollection.CountDocuments(ctx, bson.M{"tournamentid": registerTournament.TournamentId})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusOK, gin.H{"message": "error occured while checking for the team name", "hasError": true})
			return
		}

		if counted == int64(*tournament.TournamentSize) {
			c.JSON(http.StatusOK, gin.H{"message": "Registration limit reached for this tournament", "hasError": true})
			return
		}

		registerTournament.TournamentName = *tournament.Name
		registerTournament.TournamentIcon = *tournament.Icon
		registerTournament.TournamentDate = *&tournament.Date

		if tournament.Start == true {
			c.JSON(http.StatusOK, gin.H{"message": "Tournament Ongoing or finished, registration not allowed", "hasError": true})
			return
		}

		validationErr := validate.Struct(registerTournament)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message": validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		returnTournament, err := registerTournamentCollection.Find(ctx, bson.M{"tournamentid": tournament.TournamentId})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		var fil []models.RegisterTournament

		if err = returnTournament.All(ctx, &fil); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		var newData []string

		for i := range fil {
			for j := range fil[i].Players {
				newData = append(newData, fil[i].Players[j].UserName)
			}
		}

		for i := range newData {
			for j := range registerTournament.Players {
				if newData[i] == registerTournament.Players[j].UserName {
					msg := "@" + registerTournament.Players[j].UserName + " " + "has already registered for this tournament"
					c.JSON(http.StatusOK, gin.H{"message": msg, "hasError": true})
					defer cancel()
					return
				}
			}
		}

		resultInsertionNumber, insertErr := registerTournamentCollection.InsertOne(ctx, registerTournament)
		if insertErr != nil {
			msg := "item was not created"
			c.JSON(http.StatusOK, gin.H{"message": msg, "hasError": true})
			defer cancel()
			return
		}
		defer cancel()
		for j := range registerTournament.Players {
			go helper.SendEmail(registerTournament.Players[j].Email, templates.RegisterTournament(registerTournament.Players[j].UserName, registerTournament.TournamentName, registerTournament.TeamName, registerTournament.TournamentDate, registerTournament.TournamentId), "New Tournament")
		}
		c.JSON(http.StatusOK, gin.H{"counted": counted, "registerCount": tournament.TournamentSize, "message": "Registration successfull", "data": registerTournament, "hasError": false, "insertId": resultInsertionNumber})
	}
}
func GetTournament() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		// log.Fatal(id)

		// if err := helper.MatchUserTypeToUid(c, id); err != nil {
		// 	c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
		// 	return
		// }
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var tournament models.Tournament
		err := tournamentCollection.FindOne(ctx, bson.M{"tournamentid": id}).Decode(&tournament)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "tournament": tournament, "hasError": false})
	}
}

func ListPartTournament() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural": -1})

		returnTournament, err := registerTournamentCollection.Find(ctx, bson.M{"tournamentid": id}, myOptions)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		var fil []bson.M

		if err = returnTournament.All(ctx, &fil); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "tournament": fil, "hasError": false})
	}
}

func ListUserTournament() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var perPage int64 = 9
		page, err := strconv.Atoi(c.Param("page"))

		if page == 0 || page < 1 {
			page = 1
		}
		total, _ := tournamentCollection.CountDocuments(ctx, bson.M{"user_id": id})

		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural": -1})
		myOptions.SetLimit(perPage)
		myOptions.SetSkip((int64(page) - 1) * int64(perPage))

		returnTournament, err := tournamentCollection.Find(ctx, bson.M{"user_id": id}, myOptions)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		var fil []bson.M

		if err = returnTournament.All(ctx, &fil); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "total": total, "page": page, "last_page": math.Ceil(float64(total/perPage)) + 1, "tournaments": fil, "hasError": false})
	}
}

func ListUserTournamentLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		myOptions := options.Find()
		myOptions.SetLimit(7)
		myOptions.SetSort(bson.M{"$natural": -1})

		returnTournament, err := tournamentCollection.Find(ctx, bson.M{"user_id": id}, myOptions)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		var fil []bson.M

		if err = returnTournament.All(ctx, &fil); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "tournaments": fil, "hasError": false})
	}
}

func GetTournaments() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var perPage int64 = 10
		filter := bson.M{"tournamenttype": "PUBLIC"}
		page, err := strconv.Atoi(c.Param("page"))
		searchString := c.Param("search")

		if page == 0 || page < 1 {
			page = 1
		}

		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural": -1})
		myOptions.SetLimit(perPage)
		myOptions.SetSkip((int64(page) - 1) * int64(perPage))
		if searchString != "" {
			println("working")
			filter = bson.M{
				"tournamenttype": "PUBLIC",
				"$or": []bson.M{
					{
						"name": bson.M{
							"$regex": primitive.Regex{
								Pattern: searchString,
								Options: "i",
							},
						},
					},
					{
						"gamename": bson.M{
							"$regex": primitive.Regex{
								Pattern: searchString,
								Options: "i",
							},
						},
					},
				},
			}
		}
		total, _ := tournamentCollection.CountDocuments(ctx, filter)

		result, err := tournamentCollection.Find(ctx, filter, myOptions)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "error occured while listing tournaments", "hasError": true})
			return
		}
		var data []bson.M
		if err = result.All(ctx, &data); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "total": total, "page": page, "last_page": math.Ceil(float64(total/perPage)) + 1, "tournaments": data, "hasError": false})
	}
}

func GetTournamentsByMode() gin.HandlerFunc {
	return func(c *gin.Context) {
		payment := c.Param("paymentMode")
		Ttype := c.Param("tournamentType")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// fmt.Printf("Hello, %s!\n", Ttype)
		var perPage int64 = 15
		page, err := strconv.Atoi(c.Param("page"))

		if page == 0 || page < 1 {
			page = 1
		}
		total, _ := tournamentCollection.CountDocuments(ctx, bson.M{"payment": payment, "tournamenttype": "PUBLIC", "tournamentmode": Ttype})

		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural": -1})
		myOptions.SetLimit(perPage)
		myOptions.SetSkip((int64(page) - 1) * int64(perPage))
		result, err := tournamentCollection.Find(ctx, bson.M{"payment": payment, "tournamenttype": "PUBLIC", "tournamentmode": Ttype}, myOptions)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "error occured while listing tournaments", "hasError": true})
			return
		}
		var data []bson.M
		if err = result.All(ctx, &data); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "total": total, "page": page, "last_page": math.Ceil(float64(total/perPage)) + 1, "tournaments": data, "hasError": false})
	}
}

func GetTournamentsByModeLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		payment := c.Param("paymentMode")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}
		myOptions := options.Find()
		myOptions.SetLimit(10)
		myOptions.SetSort(bson.M{"$natural": -1})
		result, err := tournamentCollection.Find(ctx, bson.M{"payment": payment, "tournamenttype": "PUBLIC"}, myOptions)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "error occured while listing tournaments", "hasError": true})
			return
		}
		var data []bson.M
		if err = result.All(ctx, &data); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "tournaments": data, "hasError": false})
	}
}

// func GetTournamentByType() gin.HandlerFunc{
// 	return func(c *gin.Context) {
//         paymentType := c.Param("paymentType")

//         fmt.Println("type is", paymentType)

//         var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
//         defer cancel()

//         pipeline := []bson.M{
//             {"$match": bson.M{"payment": paymentType}},
//             {"$limit": 10}, // Limit the number of documents to 3
//         }

//         cursor, err := tournamentCollection.Aggregate(ctx, pipeline)
//         if err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true})
//             return
//         }

//         var fil []bson.M

//         if err := cursor.All(ctx, &fil); err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true})
//             return
//         }

//         c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "tournament": fil, "hasError": false})
//     }
// }

func UpdateTournament() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		primID, _ := primitive.ObjectIDFromHex(id)

		var tournament models.Tournament
		if err := c.BindJSON(&tournament); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		if err := helper.MatchUserIdToUid(c, tournament.User_id); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		filter := bson.M{"ID": primID}
		set := bson.M{"$set": bson.M{"Name": tournament.Name, "GameName": tournament.GameName, "TournamentType": tournament.TournamentType, "Shuffle": tournament.Shuffle, "Team": tournament.Team, "TournamentMode": tournament.TournamentMode, "TournamentSize": tournament.TournamentSize}}
		value, err := tournamentCollection.UpdateOne(ctx, filter, set)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Tournament updated successfully", "data": value, "tournament": id, "hasError": false})

	}
}

func DeleteTournament() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		primID, _ := primitive.ObjectIDFromHex(id)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var tournament models.Tournament
		err := tournamentCollection.FindOne(ctx, bson.M{"tournamentid": id}).Decode(&tournament)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		sentValue := false

		if tournament.IsDeleted == true {
			sentValue = false
		} else {
			sentValue = true
		}

		filter := bson.M{"ID": primID}
		set := bson.M{"$set": bson.M{"IsDeleted": sentValue}}
		value, err := tournamentCollection.UpdateOne(ctx, filter, set)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Tournament deleted successfully", "data": value, "tournament": id, "hasError": false})

	}
}

func SuspendTournament() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		primID, _ := primitive.ObjectIDFromHex(id)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var tournament models.Tournament
		err := tournamentCollection.FindOne(ctx, bson.M{"tournamentid": id}).Decode(&tournament)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		sentValue := false

		if tournament.IsSuspended == true {
			sentValue = false
		} else {
			sentValue = true
		}

		filter := bson.M{"_id": primID}
		set := bson.M{"$set": bson.M{"issuspended": sentValue}}
		value, err := tournamentCollection.UpdateOne(ctx, filter, set)
		fmt.Printf("%+v\n", value)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Tournament suspended successfully", "data": value, "tournament": id, "hasError": false})

	}
}

func UserTournaments() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var perPage int64 = 9
		page, err := strconv.Atoi(c.Param("page"))

		if page == 0 || page < 1 {
			page = 1
		}
		total, _ := registerTournamentCollection.CountDocuments(ctx, bson.M{"players.user_id": id})

		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural": -1})
		myOptions.SetLimit(perPage)
		myOptions.SetSkip((int64(page) - 1) * int64(perPage))

		result, err := registerTournamentCollection.Find(ctx, bson.M{"players.user_id": id}, myOptions)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "error occured listing user data", "hasError": true})
			return
		}

		var data []bson.M
		if err = result.All(ctx, &data); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "error occured listing data", "hasError": true})
		}

		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "total": total, "page": page, "last_page": math.Ceil(float64(total/perPage)) + 1, "tournaments": data, "hasError": false})
	}

}

func UserTournamentsLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		myOptions := options.Find()
		myOptions.SetLimit(6)
		myOptions.SetSort(bson.M{"$natural": -1})
		result, err := registerTournamentCollection.Find(ctx, bson.M{"players.user_id": id}, myOptions)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "error occured listing user data", "hasError": true})
			return
		}

		var data []bson.M
		if err = result.All(ctx, &data); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "error occured listing data", "hasError": true})
		}

		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "tournaments": data, "hasError": false})
	}

}

func RemoveUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		tournamentId := c.Param("tournamentId")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var tournament models.Tournament
		err := tournamentCollection.FindOne(ctx, bson.M{"tournamentid": tournamentId}).Decode(&tournament)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		if tournament.Start == true {
			c.JSON(http.StatusOK, gin.H{"message": "Tournament Ongoing or finished, unregistration not allowed", "hasError": true})
			return
		}

		result, err := registerTournamentCollection.Find(ctx, bson.M{"players.user_id": userId, "tournamentid": tournamentId})

		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		var data []models.RegisterTournament
		if err = result.All(ctx, &data); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		if data == nil {
			c.JSON(http.StatusOK, gin.H{"message": "you have not registered for this tournament yet", "hasError": true})
			return
		}

		var newData []models.Player
		// c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": data})
		// return

		for i, _ := range data[0].Players {
			if data[0].Players[i].User_id != userId {
				newData = append(newData, data[0].Players[i])
			}
		}

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		if newData == nil {
			err := registerTournamentCollection.FindOneAndDelete(ctx, bson.M{"registertournamentid": data[0].RegisterTournamentId}).Decode(&data[0])
			defer cancel()
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
				return
			}
		} else {
			// data[0].Players = newData
			filter := bson.M{"registertournamentid": data[0].RegisterTournamentId}
			set := bson.M{"$set": bson.M{"players": newData}}
			_, err := registerTournamentCollection.UpdateOne(ctx, filter, set)

			if err != nil {
				c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
				defer cancel()
				return
			}

		}
		c.JSON(http.StatusOK, gin.H{"message": "You have been removed from tournament", "tournaments": newData, "hasError": false})
	}

}

func inviteUser(ctx context.Context, userId, tournamentId string, user *models.User, tournament *models.Tournament) error {
	var invitedTournament models.InviteTournament

	invitedTournament.Tournament_id = tournamentId
	invitedTournament.User_id = userId
	invitedTournament.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	invitedTournament.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	if user.Email == nil || user.First_name == nil || user.Last_name == nil || tournament.Name == nil {

		return errors.New("email or firstname or lastname doesn't exist")
	}

	_, err := inviteTournamentCollection.InsertOne(ctx, invitedTournament)
	if err != nil {
		return err
	}

	var notification models.Notification

	notificationMessage := fmt.Sprintf("You have been invited to %s tournament", *tournament.Name)

	notification.UserID = userId
	notification.Message = notificationMessage

	_, err = CreateNotificationLogic(notification)
	if err != nil {
		return err
	}
	go helper.SendEmail(*user.Email, templates.TournamentInvite(*user.First_name+" "+*user.Last_name, tournamentId, *tournament.Name), "You have been Invited to a tournament")

	return nil
}

func InviteUserToTournament() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userID")
		tournamentId := c.Param("tournamentID")

		hexUserId, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid tournament id parameter", "hasError": true})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var tournament models.Tournament
		err = tournamentCollection.FindOne(ctx, bson.M{"tournamentid": tournamentId}).Decode(&tournament)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "hasError": true})
			return
		}

		var user models.User
		err = userCollection.FindOne(ctx, bson.M{"_id": hexUserId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "hasError": true})
			return
		}

		for _, invite := range tournament.AcceptedInvites {
			if invite == userId {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "User already invited to tournament", "hasError": true})
				return
			}
		}

		if err := inviteUser(ctx, userId, tournamentId, &user, &tournament); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true, "message": "Item was not created"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Invite sent successfully", "hasError": false})
	}
}

func AcceptInvite() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := c.Param("userId")
		tournamentId := c.Param("tournamentId")

		fmt.Println(userId, tournamentId)

		hexTournamentId, err := primitive.ObjectIDFromHex(tournamentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid tournament id parameter", "hasError": true})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var tournament models.Tournament
		err = tournamentCollection.FindOne(ctx, bson.M{"tournamentid": tournamentId}).Decode(&tournament)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "hasError": true})
			return
		}

		invitedUsers := tournament.AcceptedInvites

		invitedUsers = append(invitedUsers, userId)

		update := bson.D{{"$set", bson.D{{"AcceptedInvites", invitedUsers}}}}

		result, err := tournamentCollection.UpdateByID(ctx, hexTournamentId, update)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true})

			defer cancel()
			return
		}

		c.JSON(http.StatusOK, gin.H{"hasError": false, "data": result})
	}
}
