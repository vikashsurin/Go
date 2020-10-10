package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	data := []byte("hello world\n")
	/// create a file myFile and write DATA data
	// WriteFile method will create a new file with given name ,with given permissions.
	// if the file already exists it will override the data
	ioutil.WriteFile("myFile", data, 0644)

	//read File

	readFile()

}

// read a file using ioutil.
func readFile() {

	///read data from file .
	data, err := ioutil.ReadFile("myFile")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
