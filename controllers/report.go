package controllers

import (
	"context"
	// "log"
	"net/http"
	// "strconv"
	"time"

	// "github.com/Gameware/database"
	"github.com/Gameware/models"
	"github.com/gin-gonic/gin"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)


func TounamentsByDate() gin.HandlerFunc{
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

