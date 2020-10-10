package main

import (
	"fmt"
	"os"
)

//Reader ......
type Reader interface {
	Read(p []byte) (n int, err error)
}

//Writer .....
type Writer interface {
	Write(p []byte) (n int, err error)
}

func main() {
	///create a file and write data
	file, err := os.Create("myfile")
	if err != nil {
		panic(err)
	}
	//data
	data := []byte("Hello world")

	//write []bytes as a string into file
	file.Write(data)

	///write data directly as a string
	file.WriteString("This is a string")
	file.Close()
	open()
}

// OPEN file and read data
func open() {
	file, err := os.Open("myfile")
	if err != nil {
		panic(err)
	}
	//make []byte
	data := make([]byte, 30)

	/// read specific data from the file
	file.Seek(6, 0) //seek()

	//read file
	file.Read(data)

	//format data = []byte,  into string
	fmt.Printf("The file data is %v\n", string(data))

	//reset read to index
	file.Seek(0, 0)

	//make new []byte
	newData := make([]byte, 5)

	file.Read(newData)
	fmt.Println("this is new data", string(newData))
	file.Close()

}
