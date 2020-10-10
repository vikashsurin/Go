package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/post", body)
	http.HandleFunc("/get", url)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am HOMEPAGE")
}
func body(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	bs, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintln(w, bs)

	r.Body.Close()
	fmt.Fprintln(w, "I am HOMEPAGE")
}
func url(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	//list all the queries map[key]string{}
	qq := r.URL.Query()
	fmt.Printf("%+v\n", qq)
	for k, v := range qq {
		fmt.Println(k, v)
	}
	//particular query
	v := r.FormValue("q")
	os.Stdout.Write([]byte(v))

	fmt.Fprintln(w, "I am HOMEPAGE")
}
