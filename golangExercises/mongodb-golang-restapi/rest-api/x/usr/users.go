package usr

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"x/config"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type session struct {
	name         string
	lastActivity time.Time
}

var dbUsers = map[string]User{}
var dbSession = map[string]session{}

//AllUser ...
func AllUser(r *http.Request) ([]bson.M, error) {
	config.ConnectDB()
	var users []bson.M
	cursor, err := config.UserColl.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	// if err = cursor.All(context.TODO(), &users); err != nil {
	// 	log.Fatal(err)
	// }
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user bson.M
		if err = cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
		users = append(users, user)
	}
	// fmt.Printf("%+v\n", users)
	return users, nil

}

//OneUser ...
func OneUser(r *http.Request) (bson.M, error) {
	config.ConnectDB()
	// user := User{}

	var user bson.M
	err := config.UserColl.FindOne(context.TODO(), bson.M{}).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	fmt.Printf("%+v", user)
	return user, nil
}

//CurrentUser ...
func CurrentUser(r *http.Request) primitive.ObjectID {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("error")
	}
	session := dbSession[cookie.Value]
	user := dbUsers[session.name]
	fmt.Printf("%+v\n", user.ID)
	fmt.Printf("%T\n", user.ID)
	return user.ID
}

//CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) (User, error) {
	// ID:       primitive.ObjectID(),
	user := User{
		Name:  "vikash",
		Email: "vikash@gmail.com",
		// Password: "mypassword",
		Avatar: "myavatar",
		Date:   time.Now(),
	}
	p := "mypassword"

	//Check whether the email is already taken.
	opts := options.Count().SetMaxTime(2 * time.Second)
	count, err := config.UserColl.CountDocuments(context.TODO(),
		bson.D{{"email", user.Email}}, opts)
	if err != nil {
		return user, err
	}
	if count > 0 {
		fmt.Println("user  with this email already exists")
		return user, err
	}
	//create session
	sID := uuid.NewV4()
	cookie := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	cookie.MaxAge = sessionLength

	//save the cookie
	http.SetCookie(w, cookie)
	dbSession[cookie.Value] = session{user.Name, time.Now()}

	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	user.Password = bs

	//store the user
	result, err := config.UserColl.InsertOne(context.TODO(), user)
	if err != nil {
		return user, err
	}

	dbUsers[user.Name] = user
	fmt.Println("dbUsers :: ", dbUsers)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	fmt.Println(result.InsertedID)

	return user, nil
}

//UpdateUser ...
func UpdateUser(r *http.Request) (bson.M, error) {
	user := bson.M{}
	id, _ := primitive.ObjectIDFromHex("5f47603eace3ebd8e8db9657")

	result, err := config.UserColl.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		// bson.M{"name": "Jira"},
		bson.D{
			{"$set", bson.D{{"email", "king@gmail.com"}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result.ModifiedCount)
	return user, nil
}

//DeleteUser ...
func DeleteUser(r *http.Request) (bson.M, error) {
	user := bson.M{}
	id, _ := primitive.ObjectIDFromHex("5f46248ec09f14848f922e0a")

	result, err := config.UserColl.DeleteOne(
		context.TODO(),
		bson.M{"_id": id},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result.DeletedCount)
	return user, nil
}

//LoginUser ...
func LoginUser(w http.ResponseWriter, r *http.Request) (User, error) {
	user := User{}
	email := "vikash@gmail.com"
	password := "mypassword"
	// email := r.FormValue("email")
	// password := r.FormValue("password")
	err := config.UserColl.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		fmt.Println("password not matched")
		log.Fatal(err)
		return user, err
	}

	sID := uuid.NewV4()
	cookie := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	dbSession[cookie.Value] = session{user.Name, time.Now()}
	dbUsers[user.Name] = user
	fmt.Println("logged in successful.")
	return user, nil
}
