package posts

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"x/config"
	"x/usr"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreatePost ...
func CreatePost(r *http.Request) (Post, error) {
	config.ConnectDB()
	id := usr.CurrentUser(r)
	post := Post{
		User:  id,
		Title: "Computer",
		Desc:  "A computer is a powerful machine.",
		Date:  time.Now(),
	}

	result, err := config.PostColl.InsertOne(context.TODO(), post)
	if err != nil {
		fmt.Println("error create post ", err)
		return post, err

	}
	fmt.Println("Created :: ", result.InsertedID)
	return post, nil
}

//AllPosts ...
func AllPosts(r *http.Request) ([]Post, error) {
	config.ConnectDB()
	posts := []Post{}
	cursor, err := config.PostColl.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("no posts cursor found..")
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var post Post
		if err = cursor.Decode(&post); err != nil {
			log.Fatal(err)
		}
		fmt.Println(post)
		posts = append(posts, post)
	}

	return posts, nil
}

//OnePost ...
func OnePost(r *http.Request) (Post, error) {
	config.ConnectDB()
	post := Post{}
	// ObjectId("5f4e24b7396445138d4171db")
	id, _ := primitive.ObjectIDFromHex("5f4e24b7396445138d4171db")
	err := config.PostColl.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&post)
	if err != nil {
		fmt.Println("no post found ")
		return post, err
	}
	return post, nil
}

//UpdatePost ...
func UpdatePost(r *http.Request) (Post, error) {
	config.ConnectDB()
	post := Post{
		Title: "Era of computers",
		Desc:  "Computer have changed in the new era. And Now they are more efficient and powerful.",
	}
	id, _ := primitive.ObjectIDFromHex("5f4e24b7396445138d4171db")
	result, err := config.PostColl.UpdateMany(context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"title", post.Title}}},
			{"$set", bson.D{{"desc", post.Desc}}},
		},
	)
	if err != nil {
		fmt.Println("unable to update post", err)
		return post, err

	}
	fmt.Printf("%+v\n", result)
	return post, nil
}

//DeletePost ...
func DeletePost(r *http.Request) error {
	config.ConnectDB()
	id, _ := primitive.ObjectIDFromHex("5f4e24b6396445138d4171da")
	result, err := config.PostColl.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		fmt.Println("error deleting post")
		return err
	}
	fmt.Println(result.DeletedCount)
	return nil
}
