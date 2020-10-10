package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/doggy/", doggy)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I am foo")

}

func dog(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("dog.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "dog.html", nil)
}

func doggy(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "doggy.jpg")

}
