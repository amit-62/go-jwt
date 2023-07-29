package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go-jwt/helper"
	"go-jwt/models"
	"go-jwt/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection mongo.Collection = database.openCollection(database.Client, "user")
var validate = validator.New()

func SignUp() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx, cancel := context.withTimeout(context.Background(), 100*time.Second)
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

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId = c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gim.H{"error": err.error()})
			return
		}

		var ctx, cancel := context.withTimeout(context.Background, 100*time.Second)
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