package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", def)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func def(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello There!")
}

func dog(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "doggy.jpg")
}
