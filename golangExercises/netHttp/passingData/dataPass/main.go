package main

import (
	"fmt"
	"html/template"
	"io"
  "io/ioutil"
	"net/http"
  "log"
  "os"
  "path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/formPost", formPost)
	http.HandleFunc("/formGet", formGet)
	http.HandleFunc("/temp", temp)
  http.HandleFunc("/formFile",formFile)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

//url
func index(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	fmt.Fprintln(w, "Do my search: "+v)
}

//formPOST
func formPost(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("COntent-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
<form method="POST">
<input type="text" name="q">
<input type="submit">
</form>
    `+v)
}

//formGET
func formGet(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("COntent-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
<form method="GET">
<input type="text" name="q">
<input type="submit">
</form>
    `+v)
}

//from Template
func temp(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("first")
	l := r.FormValue("last")
	s := r.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(w, "index.gohtml",person{f,l,s} )
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

///formFile
func formFile(w http.ResponseWriter, r *http.Request){
    var s string

    fmt.Println(r.Method)
    if r.Method == http.MethodPost{

      // open - read - store on the server -
      f,h,err := r.FormFile("q")
      if err !=nil{
        http.Error(w,err.Error(),http.StatusInternalServerError)
        return
      }
      defer f.Close()


    fmt.Println("\n file: ",f,"\nheader: ",h, "\nerr: ",err)

    // read
    bs, err := ioutil.ReadAll(f )
    if err!=nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    s = string(bs)


    // store on the server
dst, err:=  os.Create(filepath.Join("./user/",h.Filename))
    if err !=nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    defer dst.Close()

    _,err = dst.Write(bs)
    if err!=nil{
      http.Error(w, err.Error(),http.StatusInternalServerError)
      return
    }
  }
w.Header().Set("Content-Type","text/html; charset=utf-8")
  io.WriteString(w,`
  <form method="POST" enctype="multipart/form-data">
  <input type="file" name="q"  >
  <input type="submit">
  </form>
    `+s)
}
