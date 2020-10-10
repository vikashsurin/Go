package books

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// type book struct {
// 	Isbn   int    `json:"isbn" bsno:"isbn"`
// 	Title  string `json:"title" bson:"title"`
// 	Author string `json:"author" bason:"author"`
// }

//Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	// b := book{}
	// b := Book{}
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	for k, v := range r.URL.Query() {
		fmt.Printf("%s: %s\n", k, v)
	}
	books, err := AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusMethodNotAllowed)
		return
	}
	// io.WriteString(w, books)
	fmt.Fprint(w, books)

}

//Show ...
func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(500), http.StatusMethodNotAllowed)
		return
	}
	book, err := OneBook(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, book)
	// dump(r)
}
func dump(r *http.Request) {
	output, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
		return
	}
	fmt.Println(string(output))
}

//Create ...
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	book, err := PutBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return

	}
	fmt.Fprint(w, "Created", book)
}

//Update ...
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	book, err := UpdateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}
	fmt.Fprint(w, "Book updated ", book)
}

//Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	err := DeleteBook(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "deleted")
}
