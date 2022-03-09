package controllers

import (
	"context"
	"fmt"
	"log"
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

func SaveTournament() gin.HandlerFunc{
	return func(c *gin.Context){
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
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		resultInsertionNumber, insertErr := tournamentCollection.InsertOne(ctx, tournament)
		if insertErr !=nil {
			msg := "item was not created"
			c.JSON(http.StatusOK, gin.H{"message":msg, "hasError": true})
			defer cancel()
			return
		}
		defer cancel()
		helper.SendEmail(c.GetString("email"), templates.CreateTournament(c.GetString("first_name") + " " + c.GetString("last_name")), "New Tournament")

		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "data":tournament, "hasError": false, "insertId": resultInsertionNumber})
	}
}
// userID 61e99c5efa54a7d01ff272ce
// tournamentID 61e6c5b175740deeec73b156
func RegisterTournament()gin.HandlerFunc{
	return func(c *gin.Context){
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
			c.JSON(http.StatusOK, gin.H{"message":"error occured while checking for the team name", "hasError": true})
			return
		}

		if count > 0 {
			c.JSON(http.StatusOK, gin.H{"message":"Team name is not available", "hasError": true})
			return
		}
		


		var tournament models.Tournament
		err = tournamentCollection.FindOne(ctx, bson.M{"tournamentid":tournamentId}).Decode(&tournament)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		
		registerTournament.TournamentName = *tournament.Name
		registerTournament.TournamentIcon = *tournament.Icon
		registerTournament.TournamentDate = *&tournament.Date

		if tournament.Start == true {
			if err != nil{
				c.JSON(http.StatusOK, gin.H{"message": "Tournament Ongoing or finished, registration not allowed", "hasError": true})
				return
			}
		}

		validationErr := validate.Struct(registerTournament)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		returnTournament, err := registerTournamentCollection.Find(ctx, bson.M{"tournamentid": tournament.TournamentId})
		defer cancel()
		if err != nil{
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
					c.JSON(http.StatusOK, gin.H{"message":msg, "hasError": true})
					defer cancel()
					return
				}
			}
		}


		resultInsertionNumber, insertErr := registerTournamentCollection.InsertOne(ctx, registerTournament)
		if insertErr !=nil {
			msg := "item was not created"
			c.JSON(http.StatusOK, gin.H{"message":msg, "hasError": true})
			defer cancel()
			return
		}
		defer cancel()
		for j := range registerTournament.Players {
			helper.SendEmail(registerTournament.Players[j].Email, templates.RegisterTournament(registerTournament.Players[j].UserName, registerTournament.TournamentName, registerTournament.TeamName, registerTournament.TournamentDate, registerTournament.TournamentId), "New Tournament")
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "data":registerTournament, "hasError": false, "insertId": resultInsertionNumber})
	}
}
func GetTournament() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		// log.Fatal(id)

		// if err := helper.MatchUserTypeToUid(c, id); err != nil {
		// 	c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
		// 	return
		// }
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var tournament models.Tournament
		err := tournamentCollection.FindOne(ctx, bson.M{"tournamentid":id}).Decode(&tournament)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "tournament":tournament, "hasError": false})
	}
}

func ListPartTournament() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural":-1})

		returnTournament, err := registerTournamentCollection.Find(ctx, bson.M{"tournamentid": id}, myOptions)
		defer cancel()
		if err != nil{
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
		
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "tournament":fil, "hasError": false})
	}
}

func ListUserTournament() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural":-1})

		returnTournament, err := tournamentCollection.Find(ctx, bson.M{"user_id": id}, myOptions)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		var fil []bson.M

		if err = returnTournament.All(ctx, &fil); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "tournaments":fil, "hasError": false})
	}
}

func ListUserTournamentLimit() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		myOptions := options.Find()
		myOptions.SetLimit(7)
		myOptions.SetSort(bson.M{"$natural":-1})

		returnTournament, err := tournamentCollection.Find(ctx, bson.M{"user_id": id}, myOptions)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		var fil []bson.M

		if err = returnTournament.All(ctx, &fil); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "tournaments":fil, "hasError": false})
	}
}

func GetTournaments() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage <1{
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 !=nil || page<1{
			page = 1
		}
		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural":-1})
		result,err := tournamentCollection.Find(ctx,  bson.M{}, myOptions)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing tournaments", "hasError": true})
			return
		}
		var data []bson.M
		if err = result.All(ctx, &data); err!=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "tournaments":data, "hasError": false})}
}

func GetTournamentsByMode() gin.HandlerFunc{
	return func(c *gin.Context){
		payment := c.Param("paymentMode")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage <1{
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 !=nil || page<1{
			page = 1
		}
		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural":-1})
		result,err := tournamentCollection.Find(ctx,  bson.M{"payment": payment}, myOptions)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing tournaments", "hasError": true})
			return
		}
		var data []bson.M
		if err = result.All(ctx, &data); err!=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "tournaments":data, "hasError": false})}
}

func GetTournamentsByModeLimit() gin.HandlerFunc{
	return func(c *gin.Context){
		payment := c.Param("paymentMode")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage <1{
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 !=nil || page<1{
			page = 1
		}
		myOptions := options.Find()
		myOptions.SetLimit(20)
		myOptions.SetSort(bson.M{"$natural":-1})
		result,err := tournamentCollection.Find(ctx,  bson.M{"payment": payment}, myOptions)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing tournaments", "hasError": true})
			return
		}
		var data []bson.M
		if err = result.All(ctx, &data); err!=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "tournaments":data, "hasError": false})}
}

func UpdateTournament() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		primID, _ :=primitive.ObjectIDFromHex(id)

		var tournament models.Tournament
		if err := c.BindJSON(&tournament); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		if err := helper.MatchUserIdToUid(c, tournament.User_id); err != nil {
			c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)


		filter := bson.M{"ID": primID}
		set := bson.M{"$set": bson.M{"Name": tournament.Name, "GameName": tournament.GameName, "TournamentType": tournament.TournamentType, "Shuffle": tournament.Shuffle, "Team": tournament.Team, "TournamentMode": tournament.TournamentMode, "TournamentSize": tournament.TournamentSize}}
		value, err := tournamentCollection.UpdateOne(ctx, filter, set)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "data": value, "tournament":id, "hasError": false})

	}	
}

func DeleteTournament() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		primID, _ :=primitive.ObjectIDFromHex(id)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var tournament models.Tournament
		err := tournamentCollection.FindOne(ctx, bson.M{"tournamentid":id}).Decode(&tournament)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		sentValue := false

		if tournament.IsDeleted == true {
			sentValue = false
		}else {
			sentValue = true
		}


		filter := bson.M{"ID": primID}
		set := bson.M{"$set": bson.M{"IsDeleted": sentValue}}
		value, err := tournamentCollection.UpdateOne(ctx, filter, set)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "data": value, "tournament":id, "hasError": false})

	}	
}

func SuspendTournament() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		primID, _ :=primitive.ObjectIDFromHex(id)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var tournament models.Tournament
		err := tournamentCollection.FindOne(ctx, bson.M{"tournamentid":id}).Decode(&tournament)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		sentValue := false

		if tournament.IsSuspended == true {
			sentValue = false
		}else {
			sentValue = true
		}


		filter := bson.M{"_id": primID}
		set := bson.M{"$set": bson.M{"issuspended": sentValue}}
		value, err := tournamentCollection.UpdateOne(ctx, filter, set)
		fmt.Printf("%+v\n", value)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "data": value, "tournament":id, "hasError": false})

	}	
}

func UserTournaments() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural":-1})
		result,err := registerTournamentCollection.Find(ctx,  bson.M{"players.user_id":id}, myOptions)
		
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured listing user data", "hasError": true})
			return
		}

		var data []bson.M
		if err = result.All(ctx, &data); err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured listing data", "hasError": true})
		}


		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "tournaments":data, "hasError": false})}
		
}

func UserTournamentsLimit() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		myOptions := options.Find()
		myOptions.SetLimit(7)
		myOptions.SetSort(bson.M{"$natural":-1})
		result,err := registerTournamentCollection.Find(ctx,  bson.M{"players.user_id":id}, myOptions)
		
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured listing user data", "hasError": true})
			return
		}

		var data []bson.M
		if err = result.All(ctx, &data); err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured listing data", "hasError": true})
		}


		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "tournaments":data, "hasError": false})}
		
}

func RemoveUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("userId")
		tournamentId := c.Param("tournamentId")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		
		
		result,err := registerTournamentCollection.Find(ctx,  bson.M{"players.user_id":userId, "tournamentid": tournamentId})
		
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
			return
		}

		var data []models.RegisterTournament
		if err = result.All(ctx, &data); err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
			return
		}

		if data == nil {
			c.JSON(http.StatusOK, gin.H{"message":"you have not registered for this tournament yet", "hasError": true})
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


		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
			defer cancel()
			return
		}

		if (newData == nil) {
			err := registerTournamentCollection.FindOneAndDelete(ctx, bson.M{"registertournamentid":data[0].RegisterTournamentId}).Decode(&data[0])
			defer cancel()
			if err != nil{
				c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
				return
			}
		}else {
			// data[0].Players = newData
			filter := bson.M{"registertournamentid":data[0].RegisterTournamentId}
			set := bson.M{"$set": bson.M{"players": newData}}
			_, err := registerTournamentCollection.UpdateOne(ctx, filter, set)

			if err!=nil{
				c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
				defer cancel()
				return
			}

		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "tournaments":newData, "hasError": false})}
		
}