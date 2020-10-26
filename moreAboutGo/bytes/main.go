package main

import (
	"fmt"
	"unsafe"
)
 
/*
 * []bytes are larger in size than strings
 * use string.Builder to build or modify strings.
 
*/

func main(){
var str string = "hi"
fmt.Println("string size", unsafe.Sizeof(str))
 fmt.Println("string size", str,unsafe.Sizeof(str))
slice:= []byte(str)
fmt.Println("slice size", unsafe.Sizeof(slice))

}