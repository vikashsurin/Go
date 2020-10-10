package main

/*
3 classes to types in golang.

built-int:int,string,..--------->VALUE SEMANTIC
reference types :[],maps(),channels, functions(),methods()
                  --------> VALUE SEMANTIC exception for POINTER(decode and unmarshal)
user-defined type:struct[] ---------> {WE HAVE TO CHOOSE}
  we have to choose between semantics because for the sake of CORRECTNESS.
  for example: we  should not mutate on the copy of the type user ,
			this would create confusion.
  THE DATA WHICH IS SHARED AMONG MANY OTHER FUNCTIONS OR LOCATIONS
  SHOULD USE POINTER SEMANTIC, TO reduce BUGS.

*/
func main() {

}
