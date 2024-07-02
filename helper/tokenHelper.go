package helper

import (
	// "go/token"
	"context"
	"fmt"
	// "go/token"
	"log"
	"os"
	"time"

	"github.com/PoudelAmrit123/go-rsm/database"

	"github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Clinet, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, first_Name string, lastName string, uid string) (signedToken string, signedRefreshToken string, err error) {

	claim := &SignedDetails{
		Email:      email,
		First_name: first_Name,
		Last_name:  lastName,
		Uid:        uid,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaim := &SignedDetails{

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return

	}
	return token, refreshToken, err

}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: Updated_at})

	filter := bson.M{"user_id": userId}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
	)
	defer cancel()
	if err != nil {
		log.Panic(err)
		return
	}

	return

}

func ValidateToken(signedToken string) (claim *SignedDetails, msg string) {

	//the token is invalid
	// the tokne is expired

	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	// token.Claims.(*SignedDetails)
	claim, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid ")
		msg = err.Error()
	}

	//the token is expired

	if claim.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return claim, msg

}
