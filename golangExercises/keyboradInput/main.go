package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name :")
	name, _ := reader.ReadString('\n')
	fmt.Print("hello", name)
}
