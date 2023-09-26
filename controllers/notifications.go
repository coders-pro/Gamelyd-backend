package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Gameware/database"
	"github.com/Gameware/models"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var notificationsCollection *mongo.Collection = database.OpenCollection(database.Client, "notification")

func CreateNotificationLogic(notification models.Notification) (string, error) {
	notification.Timestamp = time.Now()
	notification.IsRead = false

	_, err := notificationsCollection.InsertOne(context.TODO(), notification)

	if err != nil {
		return "", err
	}

	return "Notification Created Successfully", nil
}

func CreateNotification() gin.HandlerFunc {
	return func(c *gin.Context) {
		var notification models.Notification
		if err := c.ShouldBindJSON(&notification); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "hasError": true})
			return
		}

		msg, err := CreateNotificationLogic(notification)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true})
		}

		c.JSON(http.StatusOK, gin.H{"message": msg, "hasError": false})

	}
}

func GetNotifications() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		userID := c.Param("userID")
		cursor, err := notificationsCollection.Find(ctx, gin.H{"user_id": userID})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notifications", "hasError": true})
			defer cancel()
			return
		}

		var notifications []models.Notification
		if err := cursor.All(ctx, &notifications); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse notifications", "hasError": true})
			defer cancel()
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, gin.H{"data": notifications, "hasError": false})
	}
}

func MarkNotificationToIsRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter", "hasError": true})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		update := bson.D{{"$set", bson.D{{"is_read", true}}}}

		result, err := notificationsCollection.UpdateByID(ctx, objectID, update)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true})

			defer cancel()
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Notification updated successfully", "hasError": false, "data": result})

	}
}

func MarkAllNotificationsAsRead() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID := c.Param("userID")

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{{"user_id", userID}}
		update := bson.D{{"$set", bson.D{{"is_read", true}}}}

		_, err := notificationsCollection.UpdateMany(ctx, filter, update)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"hasError": true, "error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"hasError": false, "message": "Notifications updated successfully"})
	}
}

func DeleteNotification() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		objectID, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter", "hasError": true})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{{"_id", objectID}}

		result, err := notificationsCollection.DeleteOne(ctx, filter)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "hasError": true})
			return
		}

		c.JSON(http.StatusOK, gin.H{"hasError": false, "data": result})
	}
}
