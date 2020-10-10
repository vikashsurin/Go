package posts

import (
	"fmt"
	"net/http"
)

//Create ...
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("not Post method")
		return
	}
	post, err := CreatePost(r)
	if err != nil {
		fmt.Print("handler no post create", err)
	}
	fmt.Fprint(w, post)

}

//ShowPosts ...
func ShowPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("not get method")
		return
	}
	posts, err := AllPosts(r)
	if err != nil {
		fmt.Print("handler no post create", err)
	}
	fmt.Fprint(w, posts)

}

//ShowPost ...
func ShowPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("not get method")
		return
	}
	post, err := OnePost(r)
	if err != nil {
		fmt.Print("handler no post create", err)
	}
	fmt.Fprint(w, post)

}

//Update ...
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("not Post method in update post")
		return
	}
	post, err := UpdatePost(r)
	if err != nil {
		fmt.Print("handler no post create", err)
	}
	fmt.Fprint(w, post)

}

//Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("not Post method in update post")
		return
	}
	err := DeletePost(r)
	if err != nil {
		fmt.Print("handler no post create", err)
	}

}
