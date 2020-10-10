package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("file-exists.go"); err == nil {
		fmt.Println("The file exist in the local storage")

	} else {
		fmt.Println("File does not exist")
	}
}
