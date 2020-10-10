package main

import "net/http"

func getUser(r *http.Request) user {
	var u user
	//get Cookie
	c, err := r.Cookie("session")
	if err != nil {
		return u
	}

	//if user already exists, get user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUser[un]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {

	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUser[un]
	return ok
}
