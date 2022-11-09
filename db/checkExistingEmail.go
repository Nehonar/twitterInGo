package db

import (
	"context"
	"time"

	"github.com/nehonar/twitteringo/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*
CheckExistingEmail check in MongoDB if exists user email
*/
func CheckExistingEmail(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoDBConnection.Database("twitteringo")
	collection := db.Collection("user")

	condition := bson.M{"email": email}

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
