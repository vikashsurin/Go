package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
  "os"
	"net/http"
  "path/filepath"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var s string
	if r.Method == http.MethodPost {
		//open file
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, "file not found", http.StatusInternalServerError)
			return
		}
		defer f.Close()
		fmt.Println("\nfile: ", f, "\nheader: ", h, "\nerr: ", err)

		//read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, "file not found", http.StatusInternalServerError)
			return
		}
		s = string(bs)
		fmt.Println("The string is : ", s)

		//store on the serve
		// dst, err := os.Create("newFile")
    dst, err := os.Create(filepath.Join("./user/",h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
<form method="POST" enctype="multipart/form-data">
<input  type="file" name="q">
<input type="submit">
</form>
    `+s)

}
