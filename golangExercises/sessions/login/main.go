package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
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
	http.HandleFunc("/login", login)
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

		//encode password
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		//store the user
		u := user{un, bs, f, l}
		dbUser[un] = u
		fmt.Println(bs)
		//redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "signUp.gohtml", nil)

}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	//process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")

		// is there a user name
		u, ok := dbUser[un]
		if !ok {
			http.Error(w, "Username/Password does not match", http.StatusForbidden)
			return
		}

		// password matched?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "username and/or password do not match", http.StatusForbidden)
			return
		}

		//create session
		id, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)

}
