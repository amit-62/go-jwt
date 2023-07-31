package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/amit/go-jwt/helper"
	"github.com/amit/go-jwt/models"
	"github.com/amit/go-jwt/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection mongo.Collection = database.openCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string{
	bcrypt.GenerateFromPassword([]byte(password))
	if err != nil {
		log.Panic(err)
	}
	return string(byte)
}

func VarifyPassword(userPassword string, providedPassword string) (bool, string){
	err := bcrypt.CompareHashAndPassword(foundUser, userPassword)
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Println("password is not correct")
		check = false
	}

	return check, msg
}

func SignUp() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		if err:= c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gim.H{"error": err.error()})
			return
		} 

		validationError := validate.Struct(user)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.error()})
			return
		}

		count, err := userCollection.CountDocument(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking email"})
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		count, err := userCollection.CountDocument(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking phone"})
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"phone or email already exist"})
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ = helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertionNumber, insertError := userCollection.InsertOne(ctx, user)
		if insertError != nil {
			c.JSON(http.StatusInternalServerError, gin.H("error":"user not created"))
			return
		}
		defer cancel()
		c.JSON(http.StatusOk, resultInsertionNumber)
	}
}

func Login() gin.HandlerFunc{
	return func (c *gin.Context){
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		if err:= c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gim.H{"error": err.error()})
			return
		} 

		err := userCollection.FindOne(ctx, bson.M("email":user.Email)).Decode(foundUser)
		defer cancel()
		if err!= nil{
			c.JSON(http.StatusInternalServerError, gin.H("error":"email or password is incoreect"))
		}

		passwordIsValid, msg := VarifyPassword(*user.Password, *foundUser.Password)
		defer cancel()

		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H("email or password is wrong"))
			return
		}

		if foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H("user not found"))
			return
		}

		token , refresh_token, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.User_id)
		helper.UpdateAllToken(token, refresh_token, foundUser.User_id)

		err := userCollection.FindOne(ctx, bson.M("user_id": foundUser.User_id)).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H("error": err.error()))
			return
		}

		c.JSON(http.StatusOk, foundUser)

	}
}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId = c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gim.H{"error": err.error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background, 100*time.Second)
		defer cancel()

		var user = models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		if err!= nil {
			c.JSON(http.StatusInternalServerError, gim.H{"error": err.error()})
			return
		}

		c.JSON(http.StatusOk, user)
	}
}