package db

import (
	"context"
	"log"
	"time"

	"github.com/nehonar/twitteringo/helper"
	"github.com/nehonar/twitteringo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
InsertUser insert in MongoDB new user verificated
*/
func InsertUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoDBConnection.Database("twitteringo")
	collection := db.Collection("user")

	user.Password, _ = helper.Encrypt(user.Password)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Print("Insert in database went wrong")
		return "", false, err
	}

	objectId, _ := result.InsertedID.(primitive.ObjectID)

	return objectId.String(), true, nil
}
