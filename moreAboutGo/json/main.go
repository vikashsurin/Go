package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

type User struct {
	Username string
	Password string
	Email    string
}

type UserDb struct {
	Users []User
	Type  string
}

func main() {
	users := []User{
		{Username: "John Doe", Password: "change me", Email: "john@gmail.com"},
		{Username: "Will Smith ", Password: "borken teeth ", Email: "Will@gmail.com"},
	}

	db := UserDb{Users: users, Type: "Simple"}

	var buf = new(bytes.Buffer)

	enc := json.NewDecoder(buf)
	enc.Decode(&db)
	io.Copy(os.Stdout, buf)
}
