package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	/*
	  Stdout is a like a file.
	  os.Stdout returns ->(n len([]byte) ,err error )
	 it prints whatever it has , to the console.
	 os.Stdout.Write([]byte) -->takes in slice of bytes
	 os.Stdout.WriteString() -->takes a string, and write to file *f
	 io.WriteString(os.Stdout,"this is a string") ---> takes in os.Stdout and a value
	*/
	f := os.Stdout

	io.WriteString(f, "Hello this is using io.WriteString\n")
	f.WriteString("this is using *File.WriteString\n")
	f.Write([]byte("this is using io.Writer inteface{}\n"))

	/*
		os.Stdin
		os.Stdin returns ->(n len([]byte) ,err error )

	*/
	buff := make([]byte, 100) // make a slice of byte

	n, err := os.Stdin.Read(buff)
	fmt.Println(n) //length of bytes
	if err != nil {
		panic(err)
	}
	s := string(buff) //convert slice of byte into string and store it.
	fmt.Println(s)

}
