package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go-jwt/helper"
	"go-jwt/models"
	"go-jwt/database"
	"github.com/gin-gonic/gin"
)

var userCollection mongo.Collection = database.openCollection(database.Client, "user")
var validate = validator.New()

func getUser() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}