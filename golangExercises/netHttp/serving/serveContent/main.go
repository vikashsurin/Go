package main

import (
	"io"
	"net/http"
	"os"
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
	f, err := os.Open("doggy.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}
