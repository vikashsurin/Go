package main

import (
  "fmt"
  "net/http"
)

func main()  {
  http.HandleFunc("/",index)
  http.Handle("/favicon.ico",http.NotFoundHandler())
  http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter , r *http.Request){
  if r.URL.Path!= "/"{
    http.NotFound(w,r)
    return
  }
fmt.Println(r.URL.Path)
fmt.Fprintln(w, "go loook at your terminal")
}
