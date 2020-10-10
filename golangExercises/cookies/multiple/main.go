package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/abundance", abundance)
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello cookies!")
}
func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookies",
		Value: "some values",
	})
}

func abundance(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "liberty",
		Value: "I am liberty",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "I am general",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "I am specific",
	})
}

func read(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("my-cookies")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your Cookie #1 : ", c1)
	}
	c2, err := r.Cookie("liberty")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your Cookie #2 : ", c2)
	}
	c3, err := r.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your Cookie #3 : ", c3)
	}
	c4, err := r.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your Cookie #4 : ", c4)
	}
}
