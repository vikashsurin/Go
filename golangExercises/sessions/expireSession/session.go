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
	if s, ok := dbSessions[c.Value]; ok {
		u = dbUser[s.un]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {

	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	s := dbSessions[c.Value]
	_, ok := dbUser[s.un]
	return ok
}
