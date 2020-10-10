package say

import (
	"fmt"
	"utility/utility/bye"
)

//Hello ...
func Hello() {
	fmt.Println("said Hello.")
}

//Hi ...
func Hi() {
	fmt.Println("hi")
}

//Callme ...
func Callme() {
	bye.Bye()
}
