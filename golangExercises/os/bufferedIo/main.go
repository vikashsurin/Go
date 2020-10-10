package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//create a file ,newFile and write in it
	file, err := os.Create("newFile")
	if err != nil {
		panic(err)
	}
	//create a buffer for os file
	buff := bufio.NewWriter(file)

	//write data to buffer,
	buff.WriteString("Writing data to Buffer\n")
	// not yet written to the file , it is written to tht buffer

	//Move the Buffer data from the buffer
	//clears the data in the buffer and writes to the file.
	buff.Flush()

	//close the file
	file.Close()

	/// read
	bufRead()

	//append data
	appendData()
}

// Open and read  from the file.
func bufRead() {
	//open the file
	file, err := os.Open("newFile")
	if err != nil {
		panic(err)
	}
	// create a new reader
	reader := bufio.NewReader(file)

	/// read file byte
	data, err := reader.Peek(20)
	if err != nil {
		panic(err)
	}
	//format and print data
	fmt.Println(string(data))

	file.Close()

}

//append data to existing file
func appendData() {
	// Opens and existing file with append mode and READ WRITE	permissions.
	file, err := os.OpenFile("newFile", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	//append to file
	file.WriteString("Trying to append\n")

	//close file
	file.Close()

}
