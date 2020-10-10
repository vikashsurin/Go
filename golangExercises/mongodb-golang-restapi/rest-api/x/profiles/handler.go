package profiles

import (
	"fmt"
	"log"
	"net/http"
)

//Create ...
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("not post")
	}
	profile, err := CreateProfile(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, profile)
}

//CurrentProfile ...
func CurrentProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("not post")
	}
	profile, err := UserProfile(r)
	if err != nil {
		fmt.Println("error from current profile", err)
	}
	fmt.Fprint(w, profile)

}

//Update ...
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("not post")
	}
	profile, err := UpdateProfile(r)
	if err != nil {
		fmt.Println("error from current profile", err)
	}
	fmt.Fprint(w, profile)

}

//Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("not post")
	}
	profile, err := DeleteProfile(r)
	if err != nil {
		fmt.Println("error from current profile", err)
	}
	fmt.Fprint(w, profile, "deleted")

}
