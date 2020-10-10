package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var tpl *template.Template
var dbUser = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
	fmt.Println("user ", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	if alreadyLoggedIn(r) {
		tpl.ExecuteTemplate(w, "bar.gohtml", u)
	}

}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//process form
	if r.Method == http.MethodPost {
		// get form values
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")

		//username already taken
		if _, ok := dbUser[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		//create session
		id, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}

		//save cookie
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		//store the user
		u := user{un, p, f, l}
		dbUser[un] = u

		//redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "signUp.gohtml", nil)

}
