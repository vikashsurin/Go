package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/book", book)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Println("/")
}

var xx = map[string]string{
	"vik": "hello dear",
}

func book(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "error", 404)
	}
	// for k, v := range r.URL.Query() {
	// 	if k == "vik" {
	// 		fmt.Println("found v", v)
	// 	}
	// 	fmt.Printf("%s: %s\n", k, v)
	// }
	q := r.URL.Query()["k"]
	if q[0] == "vik" {
		fmt.Println("found")
	}
	fmt.Println(q)
	// ans := xx[q[0]]
	// fmt.Println(ans)
	// fmt.Println("/book")
	// fmt.Println("this is qxx", q)
	// fmt.Printf("%T\n", q)
}
