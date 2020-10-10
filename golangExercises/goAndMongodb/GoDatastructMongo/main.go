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
	// "gopkg.in/mgo.v2/bson"
)

// Podcast ...
type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

// Episode ...
type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

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
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fmt.Println("list error", err)
	}
	fmt.Println(databases)

	database := client.Database("quickstart")
	episodesCollection := database.Collection("episodes")
	podcastsCollection := database.Collection("podcasts")

	mongoPodcast := Podcast{
		Title:  "The MongoDB Podcast 2",
		Author: "joker",
		Tags:   []string{"mongodb", "nosql"},
	}

	insertResult, err := podcastsCollection.InsertOne(ctx, mongoPodcast)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted %v!\n", insertResult.InsertedID)

	//podcasts
	var podcasts []Podcast
	podcastCursor, err := podcastsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = podcastCursor.All(ctx, &podcasts); err != nil {
		panic(err)
	}
	fmt.Println(podcasts)

	//episodes
	var episodes []Episode
	episodeCursor, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = episodeCursor.All(ctx, &episodes); err != nil {
		panic(err)
	}
	fmt.Println(episodes)

}
