package main

import (
	"fmt"
	"os"
)

func main() {
	// by default the args array will contain , the runtime as 0 index
	fmt.Println(len(os.Args))

	//all the commands in the command line , will be collected in the os.Args ARRAY
	// Args[0] will always contain the binary executeable runtime name with the path
	fmt.Println(os.Args[1])
}
