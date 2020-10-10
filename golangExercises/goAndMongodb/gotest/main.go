package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var client *mongo.Client

//Person ...
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

// func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {}
// func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request)    {}
// func GetPersonEndpoint(response http.ResponseWriter, request *http.Request)    {}

var peopleCollection *mongo.Collection

func main() {

	fmt.Println("Starting the application...")
	ctx, cancel := context.WithTimeout(context.Background(), 3600*time.Second)
	defer cancel()
	uri := `mongodb+srv://vikash:123456vik@001.4vcgf.mongodb.net/test?retryWrites=true&w=majority`
	clientOptions := options.Client().ApplyURI(uri)
	client, _ = mongo.Connect(ctx, clientOptions)

	peopleCollection = client.Database("test").Collection("people")

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	fmt.Println(databases)

	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	fmt.Println("looking out for person : ", person)
	_ = json.NewDecoder(request.Body).Decode(&person)
	fmt.Println("after decoding : ", person)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	fmt.Printf("%T\n", ctx)
	result, _ := peopleCollection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
	fmt.Println("foun person :", person)
}

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := peopleCollection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}
func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []Person
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := peopleCollection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}
