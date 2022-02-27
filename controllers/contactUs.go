package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Gameware/database"
	"github.com/Gameware/models"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var contactCollection *mongo.Collection = database.OpenCollection(database.Client, "contact")

func SaveContactUs() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var contact models.ContactUs

		contact.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		contact.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		contact.ID = primitive.NewObjectID()
		contact.ContactId = contact.ID.Hex()
		contact.Achived = false
		contact.IsCompleted = false
		contact.IsDeleted = false

		if err := c.BindJSON(&contact); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		validationErr := validate.Struct(contact)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		resultInsertionNumber, insertErr := contactCollection.InsertOne(ctx, contact)
		if insertErr !=nil {
			c.JSON(http.StatusOK, gin.H{"message":"Error creating new entry ", "hasError": true})
			defer cancel()
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"message": "Successfull, we will get back to you soon", "data":contact, "hasError": false, "insertId": resultInsertionNumber})
	}
}

func GetAllContactUs() gin.HandlerFunc{
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
		result,err := contactCollection.Find(ctx,  bson.M{}, myOptions)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing contacts", "hasError": true})
			return
		}
		var data []bson.M
		if err = result.All(ctx, &data); err!=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "contacts":data, "hasError": false})}
}
