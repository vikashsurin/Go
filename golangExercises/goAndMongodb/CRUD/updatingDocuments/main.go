package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	uri := "mongodb+srv://vikash:123456vik@001.4vcgf.mongodb.net/<dbname>?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("client error", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3600*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("connection error", err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("ping error", err)
	}

	database := client.Database("quickstart")
	podcastsCollection := database.Collection("podcasts")

	id, _ := primitive.ObjectIDFromHex("5f2f7df6c7b98eb89a6b6e00")
	result, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"author", "Peacock"}}},
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	result, err = podcastsCollection.UpdateMany(
		ctx,
		bson.M{"title": "The Polyglot Developer Podcast"},
		bson.D{
			{"$set", bson.D{{"author", "Hell Boy"}}},
		},
	)
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	// replace
	result, err = podcastsCollection.ReplaceOne(
		ctx,
		bson.M{"author": "Hell Boy"},
		bson.M{
			"title":  "Helo Boy title",
			"author": "Hell Girl",
		},
	)
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

}
