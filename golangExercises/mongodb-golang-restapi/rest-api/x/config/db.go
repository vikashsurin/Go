package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//UserColl ...
var UserColl *mongo.Collection

//ProfileColl ....
var ProfileColl *mongo.Collection

//PostColl ...
var PostColl *mongo.Collection

//ConnectDB ....
func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Ping error.")
		log.Fatal(err)
	}

	database := client.Database("smedia")
	UserColl = database.Collection("users")
	ProfileColl = database.Collection("profiles")
	PostColl = database.Collection("posts")
	fmt.Println("connected to db")
}
