package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie-hero",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprint(w, "COOKIE WRITTEN IN YOUR BROWSER")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie-hero")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusUnavailableForLegalReasons)
		return
	}
	fmt.Fprintln(w, "YOUR-CODE :", c)
}

func auth(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("my-cookie-hero")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("cookie from auth ::", cookie)
		h.ServeHTTP(w, r)
	})
}
