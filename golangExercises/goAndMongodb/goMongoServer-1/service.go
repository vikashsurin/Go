package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//User ...
type User struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name    string             `bson:"name,omitempty" json:"name"`
	Age     int                `bson:"age,omitrmpty" json:"age"`
	Hobbies []string           `bson:"hobbies,omitempty" json:"hobbies"`
}

//ConnectToDB exported
func ConnectToDB() (client *mongo.Client) {
	uri := "mongodb+srv://vikash:123456vik@001.4vcgf.mongodb.net/<dbname>?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	fmt.Println(databases)

	database := client.Database("test")
	userCollection := database.Collection("users")

	mongoUser := User{
		Name:    "Alphert",
		Age:     23,
		Hobbies: []string{"video games", "outdoor games", "movies", "coding"},
	}

	insertResult, err := userCollection.InsertOne(ctx, mongoUser)
	if err != nil {
		panic(err)
	}
	fmt.Println("inserted : ", insertResult.InsertedID)
	fmt.Printf("%T\n", client)
	fmt.Printf("%T\n", ctx)
	return client
}
