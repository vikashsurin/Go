package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

//DB ...
var DB *mgo.Database

//BooksColl ...
var BooksColl *mgo.Collection

func init() {
	conn, err := mgo.Dial("mongodb://localhost/library")
	if err != nil {
		panic(err)
	}
	if err = conn.Ping(); err != nil {
		panic(err)
	}
	DB = conn.DB("library")
	BooksColl = DB.C("books")
	fmt.Println("connected to library DB")
}
