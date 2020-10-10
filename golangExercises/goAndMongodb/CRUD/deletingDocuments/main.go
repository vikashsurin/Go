package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
	// episodesCollection := database.Collection("episodes")
	podcastsCollection := database.Collection("podcasts")

	//DELETE ONE
	// result, err := episodesCollection.DeleteOne(ctx, bson.M{"duration": 25})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

	//DELETE MANY
	// result, err := episodesCollection.DeleteMany(ctx, bson.M{"duration": 32})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

	//Drop Podcasts
	if err = podcastsCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}

}
