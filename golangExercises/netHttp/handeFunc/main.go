package main

import (
	"io"
	"net/http"
	"os"
  "html/template"
  "log"
)


var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.Handle("/", http.FileServer(http.Dir("."))) // file Server
  http.Handle("/resources/",http.StripPrefix("/resources",http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog1", dog1)
  http.HandleFunc("/dog2",dog2)
  http.HandleFunc("/index1",index1)
  http.HandleFunc("/index2",index2)
  http.HandleFunc("/index",index)
  http.HandleFunc("/about",about)
  http.HandleFunc("/contact",contact)
  http.HandleFunc("/apply",apply)
	http.ListenAndServe(":8080", nil)
}

func def(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "hello there")
	io.WriteString(w, `<h1>Heading</h1>`)
}

//io.Copy
func dog(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("doggy.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}

// serve Content
func dog1(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("doggy.jpg")
	if err != nil {
		http.Error(w, "file not found ", 404)
		return
	}
  defer f.Close()

  fi, err :=f.Stat()
  if err !=nil{
    http.Error(w, "file not found" , 404)
    return
  }
  http.ServeContent(w, r, fi.Name(),fi.ModTime(),f)
}
func index2(w http.ResponseWriter, r *http.Request){
  f, err :=os.Open("index.html")
  if err !=nil {
    http.Error(w,"file not found",404)
    return
  }
  defer f.Close()

  fi, err := f.Stat()
  if err !=nil{
    http.Error(w, "file not found", 404)
    return
  }
  http.ServeContent(w,r, fi.Name(),fi.ModTime(),f)
}

// serve file
func dog2(w http.ResponseWriter, r *http.Request)  {
http.ServeFile(w,r,"doggy.jpg")
}

func index1(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w,r, "index.html")
}

// static

func index(w http.ResponseWriter, r *http.Request){
  err := tpl.ExecuteTemplate(w,"index.gohtml",nil)
  HandleError(w,err)
}
func about(w http.ResponseWriter, r *http.Request){
  err := tpl.ExecuteTemplate(w,"about.gohtml",nil)
  HandleError(w,err)

}
func contact(w http.ResponseWriter, r *http.Request){
  err := tpl.ExecuteTemplate(w,"contact.gohtml",nil)
  HandleError(w,err)

}
func apply(w http.ResponseWriter, r *http.Request){

if r.Method == http.MethodPost{
  err := tpl.ExecuteTemplate(w, "applyProcess.gohtml",nil)
  HandleError(w,err)
  return
}
  err := tpl.ExecuteTemplate(w,"apply.gohtml",nil)
  HandleError(w,err)
}
func HandleError(w http.ResponseWriter , err error){
  if err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Fatalln(err)
  }
}
