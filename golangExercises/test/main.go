package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)
	http.HandleFunc("/mid", authorized(mid))
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "session",
		Value: "i-am-cookie--dfdfdf-fdfd",
	}
	http.SetCookie(w, cookie)
	fmt.Fprint(w, "SET COOKIE :: ", cookie)
}

func get(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, "GET COOKIE :: ", cookie)
}

func mid(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, "GET COOKIE :: ", cookie)
	fmt.Fprint(w, "I am covered with authorization")
}

func authorized(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("COOKIE FROM AUTH ::", cookie)
		h.ServeHTTP(w, r)
	})
}
