package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index")
}

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		fmt.Println("returned")
		// fmt.Println()
		// return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
	// enableCors(&w)
	data := struct {
		Method        string
		URL           *url.URL
		Submission    map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}
	fmt.Printf("%+v\n", data)

	if r.Method != "POST" {
		fmt.Println(" not method post")
	}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("there was an error in the read ")
	}
	defer r.Body.Close()
	os.Stdout.Write(bs)

	type word struct {
		Email    string
		Password string
	}
	var x word
	err = json.Unmarshal(bs, &x)
	// b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error")
	}

	fmt.Printf("%+v\n", x)
	bb, err := json.Marshal(x)
	if err != nil {
		fmt.Println(err)
	}
	// w.Write(bb)
	fmt.Println("from test feature")
	w.Write(bb)

	os.Stdout.Write([]byte("hello"))
}
