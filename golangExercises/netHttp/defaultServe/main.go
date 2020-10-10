package main

import (
	"io"
	"net/http"
	"text/template"
)

var tpl *template.Template

func def(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "default")
}

// func dog(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	tpl.ExecuteTemplate(w, "index.html", r.Form)
// }
func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Vikash")
}

func main() {
	http.Handle("/", http.HandlerFunc(def))
	// http.HandleFunc("/dog/", dog)
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)

}
