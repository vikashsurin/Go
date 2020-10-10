package main

import "fmt"

/*
x:=23

&x  -----> adress is a data.it needs a memory to allocate itself.
  		   returns the address of the x in the memory.
		   and we can pass this address to any function ,
		   so that the func can go to this address and
		   access tha value stored in that address.
*x  ------> pointer is a variable which stores address data.
            retuns the value located at[&x] . it cannot work on its own
			it needs an address first ,in order to look for value.
			there is no pointer without a address . if we not
			pass the address , the pointer is useless.


*/

func main() {
	x := 23
	y := &x // storing the address of x.
	fmt.Println(&x, "Original")
	fmt.Println(y)

	point(y)
	val(x)
}

/*
when , passed by reference of the memory address the address of the original
data is passed , there is no new copy made.so there is no new memory allocation
thus saving the amount of memory used and  working with a single data.
but drawback is when this address is passed to multiple function or used by
multiple function ,there is race between the functions.
it creates a bugs because all the functions are trying to read and write the data.
so the state of the data changes for every functions, because are operating at once,
------(UN VERIFIED)whether, with multiple go routines or single go routine.----
WHEN THERE ARE MULTIPLE GOROUTINES MUTATING THE DATA AT THE SPEICIED ADDRESS
DATA RACE IS CREATED.This needs to be handled explicitly.
*/
/*
func (x *int){} POINTER is not enough , * only points to value of  specific address but
it doesnot determine the  TYPE of the data , and if there is no data type
we cannot work with it.
'REMEMBER GO IS ALL ABOUT TYPES'
POINTER says give the VALUE of TYPE this stored in this ADDRESS.
*/
//POINTER SEMANTIC -->copy of the address
func point(x *int) {
	fmt.Println(x, "Reference original")

}

/*
when there are many functions which takes arguments like noPoint(),
every single time a new copy of data is created in a new memory location.
which takes up physical memory .
as a result there is a degradation in performance.
benefits are: if there are any bugs in here , it will be dealing with
copy of data. Which will not affect the original data.
#mutate data in isolation
*/
//VALUE SEMANTIC ---->what we see is what we get
func val(x int) {
	fmt.Println(&x, "Copy")
}

/*
At the time of  -func main(){}- execution GO thread is created at top level (on top of  or after the operating system )
* this is called a go routine.
* and a single go routine can work on a single block of code.(memory)
*func main(){
	x:= 23
	fmt.Println(xs)    ### 1.go routine is here
	 foo(x)            ### 3.go routine is back  here
	}

	func foo(x int){     ####2. go routine is here
		fmt.Println(x)
	}

*/
