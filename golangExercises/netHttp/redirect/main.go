package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")
}

func foo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("Hello you just posted!")
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
<form method="POST">
<input type="text" name="q">
<input type="submit">
</form>
   `)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("you visited bar", r.Method, "\n", r.URL, "\n", r.Proto, "\n", r.Header, "\n", r.Body, "\n", r.ContentLength, "\n")
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
