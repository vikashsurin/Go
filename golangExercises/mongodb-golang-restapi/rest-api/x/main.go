package main

import (
	"fmt"
	"net/http"
	"x/posts"
	"x/profiles"
	"x/usr"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create/user", usr.Create)
	http.HandleFunc("/show/users", usr.Authorized(usr.ShowUsers))
	http.HandleFunc("/show/user", usr.ShowUser)
	http.HandleFunc("/show/current", usr.Authorized(usr.Current))
	http.HandleFunc("/update/user", usr.Authorized(usr.Update))
	http.HandleFunc("/delete/user", usr.Authorized(usr.Delete))
	http.HandleFunc("/login/user", usr.Login)
	http.HandleFunc("/logout", usr.Authorized(usr.Logout))
	//profiles
	http.HandleFunc("/create/profile", usr.Authorized(profiles.Create))
	// http.HandleFunc("/show/profiles", profile.ShowProfiles)
	http.HandleFunc("/show/profile", usr.Authorized(profiles.CurrentProfile))
	http.HandleFunc("/update/profile", usr.Authorized(profiles.Update))
	http.HandleFunc("/delete/profile", usr.Authorized(profiles.Delete))
	//posts
	http.HandleFunc("/create/post", usr.Authorized(posts.Create))
	http.HandleFunc("/show/posts", posts.ShowPosts)
	http.HandleFunc("/show/post", posts.ShowPost)
	http.HandleFunc("/update/post", usr.Authorized(posts.Update))
	http.HandleFunc("/delete/post", usr.Authorized(posts.Delete))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HOMEPAGE")
}
