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

	helper "github.com/Gameware/helpers"
	"github.com/Gameware/templates"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ReportAbuseCollection *mongo.Collection = database.OpenCollection(database.Client, "reportAbuse")

func SaveReportAbuse() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var report models.ReportAbuse

		report.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		report.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		report.ID = primitive.NewObjectID()
		report.ReportAbuseId = report.ID.Hex()
		report.Achived = false
		report.IsCompleted = false
		report.IsDeleted = false

		if err := c.BindJSON(&report); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
		}

		validationErr := validate.Struct(report)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}

		resultInsertionNumber, insertErr := ReportAbuseCollection.InsertOne(ctx, report)
		if insertErr !=nil {
			c.JSON(http.StatusOK, gin.H{"message":"Error creating new entry ", "hasError": true})
			defer cancel()
			return
		}
		defer cancel()
		go helper.SendEmail("madumcbobby@yahoo.com", templates.ReportAbuse(report.Name, report.Message, report.Email), "Abuse Report From Gamelyd")

		c.JSON(http.StatusOK, gin.H{"message": "Hang on and relax we will take it from here", "data":report, "hasError": false, "insertId": resultInsertionNumber})
	}
}

func GetAllReportAbuse() gin.HandlerFunc{
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
		result,err := ReportAbuseCollection.Find(ctx,  bson.M{}, myOptions)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing contacts", "hasError": true})
			return
		}
		var data []bson.M
		if err = result.All(ctx, &data); err!=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "reports":data, "hasError": false})}
}

func DeleteReportAbuse() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var report models.ReportAbuse
		report.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		err := ReportAbuseCollection.FindOne(ctx, bson.M{"reportabuseid":id}).Decode(&report)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		sentValue := false

		if report.IsDeleted == true {
			sentValue = false
		}else {
			sentValue = true
		}


		filter := bson.M{"reportabuseid": id}
		set := bson.M{"$set": bson.M{"isdeleted": sentValue, "IsDeleted": sentValue}}
		value, err := ReportAbuseCollection.UpdateOne(ctx, filter, set)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "record updated successfully", "data": value, "recordId":id, "hasError": false})

	}	
}

func AchivedReportAbuse() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var report models.ReportAbuse

		report.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		err := ReportAbuseCollection.FindOne(ctx, bson.M{"reportabuseid":id}).Decode(&report)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		sentValue := false

		if report.Achived == true {
			sentValue = false
		}else {
			sentValue = true
		}


		filter := bson.M{"reportabuseid": id}
		set := bson.M{"$set": bson.M{"Achived": sentValue, "achived": sentValue}}
		value, err := ReportAbuseCollection.UpdateOne(ctx, filter, set)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "record updated successfully", "data": value, "recordId":id, "hasError": false})

	}	
}

func CompleteReportAbuse() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var contact models.ContactUs
		contact.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		err := ReportAbuseCollection.FindOne(ctx, bson.M{"reportabuseid":id}).Decode(&contact)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		sentValue := false

		if contact.IsCompleted == true {
			sentValue = false
		}else {
			sentValue = true
		}


		filter := bson.M{"reportabuseid": id}
		set := bson.M{"$set": bson.M{"IsCompleted": sentValue, "iscompleted": sentValue}}
		value, err := ReportAbuseCollection.UpdateOne(ctx, filter, set)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "record updated successfully", "data": value, "recordId":id, "hasError": false})

	}	
}

func GetCompleteReportAbuse() gin.HandlerFunc{
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
		result,err := ReportAbuseCollection.Find(ctx,  bson.M{"IsCompleted": true}, myOptions)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing contacts", "hasError": true})
			return
		}
		var data []bson.M
		if err = result.All(ctx, &data); err!=nil{
			log.Fatal(err)
		}
		
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "reports":data, "hasError": false})

	}	
}

func GetIsDeletedReportAbuse() gin.HandlerFunc{
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
		result,err := ReportAbuseCollection.Find(ctx,  bson.M{"IsDeleted": true}, myOptions)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing contacts", "hasError": true})
			return
		}
		var data []bson.M
		if err = result.All(ctx, &data); err!=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "contacts":data, "hasError": false})

	}	
}

func GetIsAchivedReportAbuse() gin.HandlerFunc{
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
		result,err := ReportAbuseCollection.Find(ctx,  bson.M{"Achived": true}, myOptions)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing contacts", "hasError": true})
			return
		}
		var data []bson.M
		if err = result.All(ctx, &data); err!=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "contacts":data, "hasError": false})

	}	
}

func GetActiveReportAbuse() gin.HandlerFunc{
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
		result,err := ReportAbuseCollection.Find(ctx,  bson.M{"achived": false, "isdeleted": false, "iscompleted": false}, myOptions)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing contacts", "hasError": true})
			return
		}
		var data []bson.M
		if err = result.All(ctx, &data); err!=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfully", "contacts":data, "hasError": false})

	}	
}

