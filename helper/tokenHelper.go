package helper

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"jwt/database"
	"log"
	"time"
)

type SignedDetails struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Uid       string `json:"uid"`
	User_type string `json:"user_type"`
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

var SECRET_KEY string = "JfvJGOfZMd"

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{

		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Uid:       uid,
		User_type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	var refreshClaims = &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panicln(err)
		return
	}
	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)

	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}

func UpdateAllToken(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateOjb primitive.D

	updateOjb = append(updateOjb, bson.E{"token", signedToken})
	updateOjb = append(updateOjb, bson.E{"refresh_token", signedRefreshToken})

	update_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateOjb = append(updateOjb, bson.E{"update_at", update_at})

	upsert := true
	filter := bson.M{"user_id": userId}

	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set",
				updateOjb},
		},
		&opt,
	)
	defer cancel()

	if err != nil {
		log.Panicln(err)
		return
	}
	return
}
