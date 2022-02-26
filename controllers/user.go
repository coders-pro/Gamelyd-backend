package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Gameware/database"
	helper "github.com/Gameware/helpers"
	"github.com/Gameware/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err!=nil{
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string)(bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err!= nil {
		msg = fmt.Sprintf("email of password is incorrect")
		check=false
	}
	return check, msg
}

func Signup()gin.HandlerFunc{

	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "hasError": true})
			defer cancel()
			return
			
		}

		user.IsSuspended = false
		user.IsDeleted = false
		user.EmailVerified = false
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusOK, gin.H{"message":validationErr.Error(), "hasError": true})
			defer cancel()
			return
		}


		count, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusOK, gin.H{"message":"error occured while checking for the email", "hasError": true})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		count, err = userCollection.CountDocuments(ctx, bson.M{"phone":user.Phone})
		defer cancel()
		if err!= nil {
			log.Panic(err)
			c.JSON(http.StatusOK, gin.H{"message":"error occured while checking for the phone number", "hasError": true})
			return
		}

		countName, err := userCollection.CountDocuments(ctx, bson.M{"user_name":user.User_name})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusOK, gin.H{"message":"error occured while checking for the user name", "hasError": true})
			return
		}

		if countName > 0{
			c.JSON(http.StatusOK, gin.H{"message":"user name already exists", "hasError": true})
			return
		}

		if count >0{
			c.JSON(http.StatusOK, gin.H{"message":"this email or phone number already exists", "hasError": true})
			return
		}

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr !=nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusOK, gin.H{"message":msg, "hasError": true})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "data":user, "hasError": false, "insertId": resultInsertionNumber})
	}

}

func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{"message":err.Error()})
			return 
		}

		err := userCollection.FindOne(ctx, bson.M{"email":user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message":"email or password is incorrect", "hasError": true})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true{
			c.JSON(http.StatusOK, gin.H{"message": msg, "hasError": true})
			return
		}

		if foundUser.Email == nil{
			c.JSON(http.StatusOK, gin.H{"message":"user not found", "hasError": true})
			return
		}
		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.User_id)
		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		err = userCollection.FindOne(ctx, bson.M{"user_id":foundUser.User_id}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "data":foundUser, "hasError": false})
	}
}
func CheckUserName() gin.HandlerFunc{
	return func(c *gin.Context){
		userName := c.Param("name")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)


		count, err := userCollection.CountDocuments(ctx, bson.M{"user_name":userName})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusOK, gin.H{"message":"error occured while checking for the user name", "hasError": true})
			return
		}

		if count > 0 {
			c.JSON(http.StatusOK, gin.H{"message":"Username already taken", "hasError": true, "count": count, "userName": userName})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Username available", "count": count, "userName":userName, "hasError": false})
	}
}


func GetUsers() gin.HandlerFunc{
	return func(c *gin.Context){
		// if err := helper.CheckUserType(c, "ADMIN"); err != nil {
		// 	c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
		// 	return
		// }
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		myOptions := options.Find()
		myOptions.SetSort(bson.M{"$natural":-1})
		result,err := userCollection.Find(ctx,  bson.M{}, myOptions)
		
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusOK, gin.H{"message":"error occured while listing user items", "hasError": true})
			return
		}
		var allusers []bson.M
		if err = result.All(ctx, &allusers); err!=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "users":allusers, "hasError": false})}
}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "user":user, "hasError": false})
	}
}

func DeleteUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")

		// if err := helper.MatchUserTypeToUid(c, userId); err != nil {
		// 	c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
		// 	return
		// }
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := userCollection.FindOneAndDelete(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "user":user, "hasError": false})
	}
}

func UpdateUser() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		primID, _ :=primitive.ObjectIDFromHex(id)

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}

		if err := helper.MatchUserIdToUid(c, user.User_id); err != nil {
			c.JSON(http.StatusOK, gin.H{"message":err.Error(), "hasError": true})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)


		countName, err := userCollection.CountDocuments(ctx, bson.M{"user_name":user.User_name})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusOK, gin.H{"message":"error occured while checking for the user name", "hasError": true})
			return
		}

		if countName > 0{
			c.JSON(http.StatusOK, gin.H{"message":"user name already exists", "hasError": true})
			return
		}


		filter := bson.M{"ID": primID}
		set := bson.M{"$set": bson.M{"First_name": user.First_name, "Last_name": user.Last_name, "User_name": user.User_name, "Email": user.Email, "Phone": user.Phone, "Twitter": user.Twitter, "Instagram": user.Instagram, "Facebook": user.Facebook, "Linkedin": user.Linkedin, "Country": user.Country, "Location": user.Location}}
		value, err := userCollection.UpdateOne(ctx, filter, set)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"message": err.Error(), "hasError": true})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "request processed successfullt", "data": value, "tournament":id, "hasError": false})

	}	
}