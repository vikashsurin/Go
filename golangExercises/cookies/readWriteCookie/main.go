package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/write", write)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", authorized(expire))
	// http.HandlerFunc("/auth", authen(custom))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HOME")
}
func write(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "cookie",
		Value: "some value",
	})
	fmt.Fprint(w, "cookie written in browser")
}
func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "COOKIE :", c)

}

func expire(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("cookie")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func custom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, " i am auth")
}
func authorized(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("cookie from auth ::", cookie)
		// if !alreadyLoggedIn(w, r) {
		// 	//http.Error(w, "not logged in", http.StatusUnauthorized)
		// 	http.Redirect(w, r, "/", http.StatusSeeOther)
		// 	return // don't call original handler
		// }
		h.ServeHTTP(w, r)
	})
}
