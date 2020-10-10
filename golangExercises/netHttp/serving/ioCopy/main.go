package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/doggy.jpg", dogPic)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img  src="doggy.jpg"/>`)
	io.WriteString(w, "Hello there")
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("doggy.jpg")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}

	defer f.Close()

	io.Copy(w, f)

}
