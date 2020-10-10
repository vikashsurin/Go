package usr

import (
	"fmt"
	"net/http"
	"time"
)

func alreadyLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	fmt.Println("cookies from alreadyLoggedIn", cookie)
	if err != nil {
		fmt.Println("you are not logged in.")
		return false
	}
	session, ok := dbSession[cookie.Value]
	if ok {
		session.lastActivity = time.Now()
		dbSession[cookie.Value] = session
	}
	_, ok = dbUsers[session.name]
	//refresh session
	cookie.MaxAge = sessionLength
	fmt.Println("you are already logged in here is the cookie :: ", cookie)
	return ok
}
