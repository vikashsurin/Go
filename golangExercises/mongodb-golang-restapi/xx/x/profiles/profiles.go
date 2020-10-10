package profiles

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"x/config"
	"x/usr"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CreateProfile ...
func CreateProfile(r *http.Request) (Profile, error) {
	config.ConnectDB()
	userid := usr.CurrentUser(r)
	// idStr := Hex(userid)
	profile := Profile{
		User:    userid,
		Role:    "student",
		Course:  "ba",
		Hobbies: []string{"games", "study", "code"},
		Date:    time.Now(),
	}
	//check for profile duplication
	opts := options.Count().SetMaxTime(10 * time.Second)
	count, err := config.ProfileColl.CountDocuments(context.TODO(),
		bson.M{"_id": userid}, opts)
	if err != nil {
		return profile, err
	}
	if count > 0 {
		fmt.Println("profile exists")
		return profile, err
	}
	result, err := config.ProfileColl.InsertOne(context.TODO(), profile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted : ", result.InsertedID)
	return profile, nil
}

//UserProfile ...
func UserProfile(r *http.Request) (Profile, error) {
	config.ConnectDB()
	id := usr.CurrentUser(r)
	profile := Profile{}
	err := config.ProfileColl.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&profile)
	if err != nil {
		fmt.Println("error find profile", err)
		return profile, err
	}
	return profile, nil
}

//UpdateProfile ...
func UpdateProfile(r *http.Request) (Profile, error) {
	config.ConnectDB()
	// profile := Profile{}
	profile := Profile{
		Role:    "Lol",
		Course:  "computerScience",
		Hobbies: []string{"PUBG", "study", "code"},
		Date:    time.Now(),
	}
	id := usr.CurrentUser(r)
	result, err := config.ProfileColl.UpdateMany(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"role", profile.Role}}},
			{"$set", bson.D{{"course", profile.Course}}},
			{"$set", bson.D{{"hobbies", profile.Hobbies}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result.ModifiedCount)
	fmt.Printf("%+v\n", result)
	return profile, nil
}

//DeleteProfile ...
func DeleteProfile(r *http.Request) (bson.M, error) {
	config.ConnectDB()
	profile := bson.M{}
	id := usr.CurrentUser(r)

	result, err := config.ProfileColl.DeleteOne(
		context.TODO(),
		bson.M{"_id": id},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" deleted %+v\n", result.DeletedCount)
	return profile, nil
}
