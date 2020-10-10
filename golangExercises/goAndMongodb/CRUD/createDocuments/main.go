package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	ctx, _ := context.WithTimeout(context.Background(), 3600*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("connection error", err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("ping error", err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fmt.Println("list error", err)
	}
	fmt.Println(databases)

	quickStartDatabase := client.Database("quickstart")
	podcastsCollection := quickStartDatabase.Collection("podcasts")
	episodesCollection := quickStartDatabase.Collection("episodes")

	podcastsResult, err := podcastsCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Hell Boy"},
		{Key: "tags", Value: bson.A{"development", "coding", "programming"}},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcastsResult.InsertedID)

	episodesResult, err := episodesCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{Key: "podcast", Value: podcastsResult.InsertedID},
			{Key: "title", Value: "Episode #2"},
			{Key: "description", Value: "This is the second episode."},
			{Key: "duration", Value: 25},
		},
		bson.D{
			{Key: "podcast", Value: podcastsResult.InsertedID},
			{Key: "title", Value: "Episode #1"},
			{Key: "description", Value: "This is the first episode."},
			{Key: "duration", Value: 32},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(episodesResult.InsertedIDs)
}
