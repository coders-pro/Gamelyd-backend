package controllers

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"

	"time"

	"github.com/Gameware/database"
	"github.com/Gameware/models"

	helper "github.com/Gameware/helpers"
	"github.com/Gameware/templates"
	"github.com/gin-gonic/gin"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var drawCollection *mongo.Collection = database.OpenCollection(database.Client, "draw")



func Draw() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var draw models.Draw

		if err := c.BindJSON(&draw); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		draw.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		draw.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		draw.ID = primitive.NewObjectID()
		draw.DrawId = draw.ID.Hex()

		filter := bson.M{"tournamentid": draw.TournamentId }

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
		if result.Err() != nil {
			c.JSON(http.StatusOK, gin.H{"message": "Error please try again", "hasError": true})
			defer cancel()
			return
		}

		if draw.Stage == 1 {
			participants, err := registerTournamentCollection.Find(ctx, bson.M{"tournamentid": draw.TournamentId})
			defer cancel()
			if err != nil{
				c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
				defer cancel()
				return
			}
			fmt.Printf("%+v\n", "stage is 1")

			var fil []models.RegisterTournament

			if err := participants.All(ctx, &fil); err != nil {
				c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
				return
			}

			for i := range fil {
				j := rand.Intn(i + 1)
				fil[i], fil[j] = fil[j], fil[i]
			}	
			

			
			Team1 := models.Teams{
				TeamName: "",
				Players: nil,
			}
			Team2 := models.Teams{
				TeamName: "",
				Players: nil,
			}
			fmt.Println(reflect.TypeOf(Team1))
			
			count := 1
			var allData []interface{}
			var formatData []models.Draw
			
			
			for count <= len(fil) {
					if count%2 == 0 {
						Team1.Players = fil[count - 2].Players
						Team1.TeamName = fil[count - 2].TeamName
		
						Team2.Players = fil[count - 1].Players
						Team2.TeamName = fil[count - 1].TeamName
		
						draw.Team1 = Team1
						draw.Team2 = Team2
						draw.ID = primitive.NewObjectID()
						draw.Stage = 1
						draw.DrawId = draw.ID.Hex()
						formatData = append(formatData, draw)
						for i, _ := range draw.Team1.Players {
							go helper.SendEmail(draw.Team1.Players[i].Email , templates.DrawTournament(draw.Team1.Players[i].UserName, "", draw.TournamentId), "Tournament Draw")
						}
						for i, _ := range draw.Team2.Players {
							go helper.SendEmail(draw.Team2.Players[i].Email , templates.DrawTournament(draw.Team2.Players[i].UserName, "", draw.TournamentId), "Tournament Draw")
						}
						count++
		
					}else {
						count++
					}	
				
			}

			if len(fil)%2 != 0 {
				Team1.Players = fil[len(fil) - 1].Players
				Team1.TeamName = fil[len(fil) - 1].TeamName

				Team2.TeamName = "Automatic Qualification"
				Team2.Players = nil
				draw.Stage = 1
				draw.Team1 = Team1
				draw.Team2 = Team2
				draw.Winner = "Team1"
				draw.ID = primitive.NewObjectID()
				draw.DrawId = draw.ID.Hex()
				formatData = append(formatData, draw)
				for i, _ := range draw.Team1.Players {
					go helper.SendEmail(draw.Team1.Players[i].Email , templates.DrawTournament(draw.Team1.Players[i].UserName, "", draw.TournamentId), "Tournament Draw")
				}
			}
			fmt.Println(reflect.TypeOf(allData))
			for _, t := range formatData {
				allData = append(allData, t)
			}
						

			resultInsertionNumber, insertErr := drawCollection.InsertMany(ctx, allData)
			if insertErr !=nil {
				c.JSON(http.StatusOK, gin.H{"message":  insertErr, "hasError": true})
				defer cancel()
				return
			}
			fmt.Printf("%v", resultInsertionNumber)
			fmt.Printf("%+v\n", insertErr)
			filter := bson.M{"tournamentid": draw.TournamentId}
			set := bson.M{"$set": bson.M{"Start": true}}
			value, err := tournamentCollection.UpdateOne(ctx, filter, set)
			defer cancel()
			fmt.Print(value)

			
			c.JSON(http.StatusOK, gin.H{"message": "request processed successfull", "data": allData, "hasError": false, "insertIds": resultInsertionNumber})
			defer cancel()
			return
		}else {
			returnDraw, err := drawCollection.Find(ctx, bson.M{"stage": draw.Stage - 1, "tournamentid": draw.TournamentId})
			defer cancel()
			if err != nil{
				c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
				defer cancel()
				return
			}

			var fil []models.Draw
			var newDraw []models.Teams
			var request models.Draw
			var request2 models.Draw
			var submitData []models.Draw
			if err = returnDraw.All(ctx, &fil); err != nil {
				c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
				defer cancel()
				return
			}
		for i := range fil {
			if fil[i].Winner == "Team1" {
				newDraw = append(newDraw, fil[i].Team1)
			}else if fil[i].Winner == "Team2" {
				newDraw = append(newDraw, fil[i].Team2)
			}	
		}
			if len(newDraw) == 1 {
				c.JSON(http.StatusOK, gin.H{"message": "You can't draw with just one team", "hasError": true, "new": newDraw})
				defer cancel()
				return
			}
			
		if len(newDraw)%2 != 0 {
			var temp models.Teams = newDraw[0]
			newDraw[0] = newDraw[len(newDraw) - 1]
			newDraw[len(newDraw) - 1] = temp
		}
		newCount := 1

		if len(newDraw) == 2 {
			request.Team1 = newDraw[0]
			request.Team2 = newDraw[1]	
			
			request.Stage = draw.Stage
			request.ID = primitive.NewObjectID()
			request.TournamentId = draw.TournamentId
			request.DrawId = request.ID.Hex()
			request.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			request.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			for i, _ := range request.Team1.Players {
				go helper.SendEmail(request.Team1.Players[i].Email , templates.DrawTournament(request.Team1.Players[i].UserName, "", request.TournamentId), "Tournament Draw")
			}
			for i, _ := range request.Team2.Players {
				go helper.SendEmail(request.Team2.Players[i].Email , templates.DrawTournament(request.Team2.Players[i].UserName, "", request.TournamentId), "Tournament Draw")
			}
			submitData = append(submitData, request)			
		}else {
			for newCount < len(newDraw) {
				if newCount%2 == 0 {
						request.Team1 = newDraw[newCount - 2]
						request.Team2 = newDraw[newCount - 1]	
						
						request.Stage = draw.Stage
						request.ID = primitive.NewObjectID()
						request.TournamentId = draw.TournamentId
						request.DrawId = request.ID.Hex()
						request.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
						request.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
						submitData = append(submitData, request)
						for i, _ := range request.Team1.Players {
							go helper.SendEmail(request.Team1.Players[i].Email , templates.DrawTournament(request.Team1.Players[i].UserName, "", request.TournamentId), "Tournament Draw")
						}
						for i, _ := range request.Team2.Players {
							go helper.SendEmail(request.Team2.Players[i].Email , templates.DrawTournament(request.Team2.Players[i].UserName, "", request.TournamentId), "Tournament Draw")
						}									
						newCount++
				
				}else {
					newCount++
				}
		}	
	}
	if len(newDraw)%2 != 0 {
		if len(newDraw) != 2 {
			request2.Team1 =  newDraw[len(newDraw) - 1]
			request2.Team2.Players = nil
			request2.Winner = "Team1"
			request2.Stage = draw.Stage
			request2.Team2.TeamName = "Automatic Qualification"
			request2.TournamentId = draw.TournamentId
			request2.ID = primitive.NewObjectID()
			request2.DrawId = request2.ID.Hex()
			request2.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			request2.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			for i, _ := range request.Team1.Players {
				go helper.SendEmail(request.Team1.Players[i].Email , templates.DrawTournament(request.Team1.Players[i].UserName, "", request.TournamentId), "Tournament Draw")
			}

			submitData = append(submitData, request2)
		}
		
	}

		var newAll []interface{}

		for _, t := range submitData {
			newAll = append(newAll, t)
		}

		_, insertErr := drawCollection.InsertMany(ctx, newAll)
			if insertErr !=nil {
				c.JSON(http.StatusOK, gin.H{"message":  insertErr.Error(), "hasError": true})
				defer cancel()
				return
			}
			filter := bson.M{"tournamentid": draw.TournamentId}
			set := bson.M{"$set": bson.M{"Start": true}}
			value, err := tournamentCollection.UpdateOne(ctx, filter, set)
			defer cancel()
			fmt.Print(value)
	
	c.JSON(http.StatusOK, gin.H{"message": "request processed successfull", "hasError": false, "data": newAll, "newDraw": newDraw})
	defer cancel()
	return

		}
	

	}
}



func DrawBr() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var draw models.Draw

		if err := c.BindJSON(&draw); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		draw.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		draw.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		draw.ID = primitive.NewObjectID()
		draw.DrawId = draw.ID.Hex()

		returnDraw, err := drawCollection.Find(ctx, bson.M{"tournamentid": draw.TournamentId})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		var data []models.Draw

		if err = returnDraw.All(ctx, &data); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		if len(data[0].BRTeams) > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "Tournament already drawn, you just have to update your scores", "hasError": true})
			defer cancel()
			return
		}

		participants, err := registerTournamentCollection.Find(ctx, bson.M{"tournamentid": draw.TournamentId})
			defer cancel()
			if err != nil{
				c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
				defer cancel()
				return
			}
			fmt.Printf("%+v\n", "stage is 1")

			var fil []models.RegisterTournament

			if err := participants.All(ctx, &fil); err != nil {
				c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
				return
			}	
			

			var teams models.BRTeams
			
			for i := range fil {
				teams.Kills = 0
				teams.Wins = 0
				teams.FirstBloods = 0
				teams.TeamName = fil[i].TeamName
				teams.Players = fil[i].Players
				draw.BRTeams = append(draw.BRTeams, teams)
			}
			
		
			resultInsertionNumber, insertErr := drawCollection.InsertOne(ctx, draw)
			if insertErr !=nil {
				c.JSON(http.StatusOK, gin.H{"message":  insertErr, "hasError": true})
				defer cancel()
				return
			}
			fmt.Printf("%v", resultInsertionNumber)
			fmt.Printf("%+v\n", insertErr)
			filter := bson.M{"tournamentid": draw.TournamentId}
			set := bson.M{"$set": bson.M{"Start": true}}
			value, err := tournamentCollection.UpdateOne(ctx, filter, set)
			defer cancel()
			fmt.Print(value)
			
			
			c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "resultId":resultInsertionNumber,  "data": draw, "hasError": false})
			defer cancel()
			return
	}
}


func GetDrawByTornamentID() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("tornamentId")
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		returnDraw, err := drawCollection.Find(ctx, bson.M{"tournamentid": id})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		var fil []bson.M

		if err = returnDraw.All(ctx, &fil); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "draws":fil, "hasError": false})
	}
}

func AddWinner() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("drawId")
		type Winner struct {
			Winner string		`json:"Winner" validate:"required"`
		}
		var winner Winner
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&winner); err != nil {
			c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
			defer cancel()
			return
		}
		
		validationErr := validate.Struct(winner)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}
		fmt.Printf("%+v\n", ctx)
		

		filter := bson.M{"drawid": id}

		update := bson.M{
			"$set": bson.M{"winner": winner.Winner},
		}

		upsert := true
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}

		result := drawCollection.FindOneAndUpdate(ctx, filter, update, &opt)
		if result.Err() != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "draws":result, "hasError": false})
	}
}

func AddTime() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("drawId")
		type Data struct {
			Time string		`json:"Time" validate:"required"`
			Date string		`json:"Date" validate:"required"`
		}
		var data Data
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
			defer cancel()
			return
		}
		
		validationErr := validate.Struct(data)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}
		fmt.Printf("%+v\n", ctx)
		

		filter := bson.M{"drawid": id}

		update := bson.M{
			"$set": bson.M{"date": data.Date, "time": data.Time},
		}

		upsert := true
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}

		result := drawCollection.FindOneAndUpdate(ctx, filter, update, &opt)
		if result.Err() != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		var draws models.Draw 

		err := drawCollection.FindOne(ctx, bson.M{"drawid":id}).Decode(&draws)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}
		for i, _ := range draws.Team1.Players {
			go helper.SendEmail(draws.Team1.Players[i].Email , templates.AddTime(draws.Team1.Players[i].UserName, "", draws.TournamentId, data.Time, data.Date), "Tournament Time and Date")
		}
		for i, _ := range draws.Team2.Players {
			go helper.SendEmail(draws.Team2.Players[i].Email , templates.AddTime(draws.Team2.Players[i].UserName, "", draws.TournamentId, data.Time, data.Date), "Tournament Time and Date")
		}
		
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "draws":draws, "hasError": false})
	}
}

func AddScore() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("drawId")
		type Score struct {
			Team1 	interface{}			`json:"Team1" validate:"required"`
			Team2 	interface{}			`json:"Team2" validate:"required"`
			Winner 	string		`json:"Winner" validate:"required"`
		}
		var data Score
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
			defer cancel()
			return
		}
		
		validationErr := validate.Struct(data)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		filter := bson.M{"drawid": id}

		update := bson.M{
			"$set": bson.M{"Team1Score": data.Team1 , "Team2Score": data.Team2, "Winner": data.Winner, "winner": data.Winner},
		}

		upsert := true
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}

		result := drawCollection.FindOneAndUpdate(ctx, filter, update, &opt)
		if result.Err() != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "draws":result, "hasError": false})
	}
}

func AddLink() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("drawId")
		type Link struct {
			Link 	string			`json:"Link" validate:"required"`
		}
		var data Link
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
			defer cancel()
			return
		}
		
		validationErr := validate.Struct(data)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		filter := bson.M{"drawid": id}

		update := bson.M{
			"$set": bson.M{"Link": data.Link},
		}

		upsert := true
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}

		result := drawCollection.FindOneAndUpdate(ctx, filter, update, &opt)
		if result.Err() != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		var draws models.Draw 

		err := drawCollection.FindOne(ctx, bson.M{"drawid":id}).Decode(&draws)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}
		for i, _ := range draws.Team1.Players {
			go helper.SendEmail(draws.Team1.Players[i].Email , templates.AddLink(draws.Team1.Players[i].UserName, "", draws.TournamentId), "Link to join match is now available")
		}
		for i, _ := range draws.Team2.Players {
			go helper.SendEmail(draws.Team2.Players[i].Email , templates.AddLink(draws.Team2.Players[i].UserName, "", draws.TournamentId), "Link to join match is now available")
		}
		
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "draws":result, "hasError": false})
	}
}

func UpdateBrDraw() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("drawId")
		
		type Back struct {
			Players []models.BRTeams
			Link string
			Date string
			Time string
		}
		
		var data Back
		fmt.Print(data)
		var teams []models.BRTeams
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
			defer cancel()
			return
		}
		fmt.Print(data)
		validationErr := validate.Struct(data)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		for i := range data.Players {
			teams = append(teams, data.Players[i])
		}


		filter := bson.M{"drawid": id}

		update := bson.M{
			"$set": bson.M{"brteams": teams, "time": data.Time, "link": data.Link, "date": data.Date},
		}

		upsert := true
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}

		result := drawCollection.FindOneAndUpdate(ctx, filter, update, &opt)
		if result.Err() != nil {
			c.JSON(http.StatusOK, gin.H{"message":"error updating record", "hasError": true})
			defer cancel()
			return
		}

		// if data.Time != "" {

		// }
		
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "id": id, "draws":ctx, "hasError": false})
		var draws models.Draw

		err := drawCollection.FindOne(ctx, bson.M{"drawid":id}).Decode(&draws)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}
		if data.Link != "" {
			for i, _ := range data.Players {
				for j, _ := range data.Players[i].Players {
					go helper.SendEmail(data.Players[i].Players[j].Email, templates.AddLink(data.Players[i].Players[j].UserName, "", draws.TournamentId), "Link to join match is now available")
				}
			}
		}
		if data.Time != "" {
			for i, _ := range data.Players {
				for j, _ := range data.Players[i].Players {
					go helper.SendEmail(data.Players[i].Players[j].Email, templates.AddTime(data.Players[i].Players[j].UserName, "", draws.TournamentId, data.Time, data.Date), "Tournament Time and Date")
				}
			}
		}
		
		return
	}
}