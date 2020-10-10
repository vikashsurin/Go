package usr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"x/config"
)

const sessionLength int = 24 * 3600 * 1000

func cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
}

//ShowUsers ...
func ShowUsers(w http.ResponseWriter, r *http.Request) {
	config.ConnectDB()
	if r.Method != "GET" {
		fmt.Println("not get method")
	}
	users, err := AllUser(r)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, users)
	fmt.Printf("%T\n", users)
}

//ShowUser ...
func ShowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("not get method")
	}
	user, err := OneUser(r)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, user)
}

//Current ...
func Current(w http.ResponseWriter, r *http.Request) {
	CurrentUser(r)
}

//Create ...
func Create(w http.ResponseWriter, r *http.Request) {
	cors(w)
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	config.ConnectDB()
	if r.Method != "POST" {
		fmt.Println("not post")
		return
	}
	user, err := CreateUser(w, r)
	if err != nil {
		panic(err)
	}
	bs, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Write(bs)
	fmt.Fprint(w, user)

}

//Update ...
func Update(w http.ResponseWriter, r *http.Request) {
	config.ConnectDB()
	if r.Method != "POST" {
		fmt.Println("not post")
	}
	user, err := UpdateUser(r)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, user)

}

//Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	config.ConnectDB()
	if r.Method != "POST" {
		fmt.Println("not post")
	}
	user, err := DeleteUser(r)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, user)

}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	cors(w)
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	config.ConnectDB()
	if r.Method != "POST" {
		fmt.Println("not POST method")
		return
	}
	user, err := LoginUser(w, r)
	if err != nil {
		panic(err)
	}
	// fmt.Fprint(w, user, "hello")
	// io.WriteString(w, "hello there")
	// os.Stdout.Write([]byte("hello "))
	bs, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error marshal")
	}
	w.Write(bs)

}

//Logout ...
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")

	//delete session
	delete(dbSession, cookie.Value)

	//remove the cookie
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	fmt.Println("You are logged out..")
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("GET COOKIE :: ", cookie)
	}
	return cookie
}

//Authorized ...
func Authorized(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := getCookie(w, r)
		fmt.Println("cookiessssssss", c)
		cookie, err := r.Cookie("session")
		if err != nil {
			fmt.Println("NO COOKIE FOUND IN AUTH...")
			return
		}
		if !alreadyLoggedIn(r) {
			//http.Error(w, "not logged in", http.StatusUnauthorized)
			// http.Redirect(w, r, "/", http.StatusSeeOther)
			return // don't call original handler
		}
		fmt.Println("COOKIE FROM AUTH :: ", cookie)
		h.ServeHTTP(w, r)
	})
}
