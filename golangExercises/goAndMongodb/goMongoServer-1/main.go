package main

import "github.com/julienschmidt/httprouter"

func main() {

	ConnectToDB()
	ctx := ConnectToDB()

	router := httprouter.New()
}
