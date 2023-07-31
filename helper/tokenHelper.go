package helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/amit/go-jwt/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignedDetails struct {
	Email			string
	First_name		string
	Last_name		string
	Uid 			string
	User_type 		string
	jwt.StandarClaims
}

var SECRET_KEY string= os.Getenv("SECRET_KEY")


func GenerateAllTokens (email string, firstName string, lastName string, userType string, uid string)
	(signedToken string, signedRefreshToken string, err) {
		claims := &SignedDetails{
			Email : email,
			First_name : firstName,
			Last_name : lastName,
			Uid : uid,
			User_type : userType,
			StandarClaims : jwt.StandarClaims{
				ExpiresAt : time.Now().Local().Add(time.Hours*time.Duration(24)).unix(),
			}
		}

		refreshClaims := &SignedDetails{
			StandarClaims : jwt.StandarClaims{
				ExpiresAt : time.Now().Local().Add(time.Hours*time.Duration(168)).unix(),
			}
		}

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
		refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

		if err!= nil {
			log.Panic(err)
			return
		}

		return token, refreshToken, err
	}

func UpdateAllToken(signedToken string, signedRefreshToken string, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var updateObject primitive.Database

	updateObject.append(updateObject, bson.E("token":token))
	updateObject.append(updateObject, bson.E("refresh_token":refresh_token)

	Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObject.append(updateObject, bson.E("updated_at", Updated_at))

	upsert := true
	filter := bson.M("user_id":userIdser)
	obj := options.UpdateOptions{
		Upsert : &upsert
	}

	_, err := userCollection.UpdateOne{
		ctx,
		filter,
		bson.D{
			{"$set", updateObject}
		},
		&opj
	}

	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}
	return
}
	