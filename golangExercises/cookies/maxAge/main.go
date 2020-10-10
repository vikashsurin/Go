package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
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

func delete(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("my-cookie-hero")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	c.MaxAge = -1 //deletes a cookie
	http.SetCookie(w, c)
	fmt.Println(c)

}
