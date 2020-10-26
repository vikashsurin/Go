package main

import (
	"fmt"
	"unsafe"
)

/*
 *  the total size of the struct is 16 where x = 8 and y = 8
 *  which means the structure doesnot takes any space. in most sceneraios.
  * structs create empty padding for int
  * struct is of size 0, even if there are millions of empty structs nested inside the size is still 0.
  * make([]struct{},100) slice os empty struct takes memory.
  * empty structs can be used in channels for efficiency.
 */
func main() {
	var data struct {
		x uint64
		y uint64
	}
	fmt.Println(
		unsafe.Sizeof(data.x),
		"+",
		unsafe.Sizeof(data.y),
		"=",
		unsafe.Sizeof(data),
	)
}
