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
	"github.com/gin-gonic/gin"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var drawCollection *mongo.Collection = database.OpenCollection(database.Client, "draw")


func Draw() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var draw models.Draw

		if err := c.BindJSON(&draw); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		draw.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		draw.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		draw.ID = primitive.NewObjectID()
		draw.DrawId = draw.ID.Hex()


		if draw.Stage == 1 {
			participants, err := registerTournamentCollection.Find(ctx, bson.M{"tournamentid": draw.TournamentId})
			defer cancel()
			if err != nil{
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true})
				defer cancel()
				return
			}
			fmt.Printf("%+v\n", "stage is 1")

			var fil []models.RegisterTournament

			if err := participants.All(ctx, &fil); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true})
				return
			}

			if draw.Stage == 1 {
			for i := range fil {
				j := rand.Intn(i + 1)
				fil[i], fil[j] = fil[j], fil[i]
				}	
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
						draw.DrawId = draw.ID.Hex()
						formatData = append(formatData, draw)
											
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
				draw.Team1 = Team1
				draw.Team2 = Team2
				draw.Winner = "Team1"
				draw.ID = primitive.NewObjectID()
				draw.DrawId = draw.ID.Hex()

				formatData = append(formatData, draw)
			}
			fmt.Println(reflect.TypeOf(allData))
			for _, t := range formatData {
				allData = append(allData, t)
			}
		
			// if err != nil {
			// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "new error", "hasError": true})
			// 	return
			// }

		
			// resultInsertionNumber, insertErr := drawCollection.InsertMany(ctx, allData)
			// if insertErr !=nil {
			// 	c.JSON(http.StatusInternalServerError, gin.H{"error":  insertErr, "hasError": true})
			// 	defer cancel()
			// 	return
			// }
			// fmt.Printf("%v", resultInsertionNumber)
			// fmt.Printf("%+v\n", insertErr)
			
			
			c.JSON(http.StatusOK, gin.H{"message": "request processed successfull", "data": allData, "hasError": false, "new": fil})
		}else {
			returnDraw, err := drawCollection.Find(ctx, bson.M{"stage": 2, "tournamentid": draw.TournamentId})
			defer cancel()
			if err != nil{
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true})
				defer cancel()
				return
			}

			var fil []models.Draw
			// var newDraw models.Draw

			if err = returnDraw.All(ctx, &fil); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true})
				defer cancel()
				return
			}
			

			

			// for i := range fil {
			// 	if fil[i].Winner == "Team1" {
			// 		newDraw.
			// 	}
			// }

			

			
		
			// if err != nil {
			// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "new error", "hasError": true})
			// 	return
			// }

		
			// resultInsertionNumber, insertErr := drawCollection.InsertMany(ctx, allData)
			// if insertErr !=nil {
			// 	c.JSON(http.StatusInternalServerError, gin.H{"error":  insertErr, "hasError": true})
			// 	defer cancel()
			// 	return
			// }
			// fmt.Printf("%v", resultInsertionNumber)
			// fmt.Printf("%+v\n", insertErr)
			
			
			c.JSON(http.StatusOK, gin.H{"message": "request processed successfull", "hasError": false, "new": fil})
		}

	}
}


func GetDrawByTornamentID() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("tornamentId")
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// returnDraw, err := drawCollection.Find(ctx, bson.M{"tournamentid": id})
		returnDraw, err := drawCollection.Find(ctx, bson.M{"stage": 2, "tournamentid": id})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		var fil []bson.M

		if err = returnDraw.All(ctx, &fil); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true})
			defer cancel()
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "draws":fil, "hasError": false})
	}
}