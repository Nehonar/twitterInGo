package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
MongoDBConnection this is a function to connect in MongoDB
userMongoDB get env username MongoDB
passMongoDB get env password MongoDB
urlMongoDB get env url MongoDB
mongoDB is a complete url with params to connection at MongoDB
StatusConnectionOK is just a return if connections is OK with int 1
StatusConnectionWrong is just a return if connections is wrong with int 0
*/
var MongoDBConnection = ConnectionMongoDB()
var userMongoDB = os.Getenv("USERNAME_MONGODB")
var passMongoDB = os.Getenv("PASSWORD_MONGODB")
var urlMongoDB = os.Getenv("HOST_MONGODB")
var mongoDB = fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", userMongoDB, passMongoDB, urlMongoDB)
var StatusConnectionOK = 1
var StatusConnectionWrong = 0

/*
ConnectionMongoDB is to connect with MongoDB
*/
func ConnectionMongoDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDB))

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connection to MongoDB: OK")
	return client
}

/*
CheckConnectionWithPing is just a check, return int 0 if is wrong or int 1 if is OK
*/
func CheckConnectionWithPing() int {
	err := MongoDBConnection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return StatusConnectionWrong
	}
	return StatusConnectionOK
}
