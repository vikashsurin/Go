### Struct
- collection of properties that are related together.

### Updating Struct
- There are several ways to update a struct 
```
 //example method 1
 type person struct{
     firstname string
 }

 func main(){
     p1:= person{
         firstname:"Rob Pike ",
     }

     //update
     p1.firstname = "Jacob Peterson"
 }
 ```
- updating via function
- updating via function is a gotcha. Its not like traditional update via arguments .
- updating struct via function needs to use the __pointer__ spec . as go lang is pass by value.
- The context pass by value means , when a struct is passed intto the function as argument it is going to create a copy of the struct and modify or update that copy. The original struct remains the same. so the use of pointers is necessary . 
- [see example](https://play.golang.org/p/VKGmERwmNch) 
```
//example method 2
package main

import (
	"fmt"
)


type person struct {
	firstname string
}


func main() {

	p1 := person{
		firstname: "Adam",
	}
	
	fmt.Println(p1)
	
	p1.up("alex")
	fmt.Println(p1)

}

func (p *person) up(newname string) {
	(*p).firstname = newname
}

```

### Embedding Struct
- 

### Struct with Receiver functions
- 

### Pointers to update struct
- __&_variable__ gives us the memmory address of the variable .
- __*_pointer__ gives us the value , of that memory address.
- __*_type__ gives us the pointer of the type.

### Pointer Shortcut
- 

### Pointers Gotchas
- 
