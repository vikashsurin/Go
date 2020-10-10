
### How do we run go code?
  ``` 
       go run main.go
  ```
 - go run - compiles and executes the file
 - go build - compiles a bunch of go source code files and returns a executeable file
 - go fmt - format the go code
 - go get - download raw source code
 - go test - run any test files associted with the current project
 - go install - compiles and installs a  pacakge

### What does  package main mean?
- package == project == workspace
- package main ,i.e. a project belong to the package main
- the package __main__ ,the word main means , the file is to be build in future . the package with no __main__   cannot be built and cannot be made executeable.
- package __main__ can be built, package __apple__ cannot be built.

### Types of files
- executeable file (binary file)
- reuseable file (functions,vars...)

###  What does import fmt means?
```
    import "fmt"
```
- give the package __main__ access to  all the code of the package __fmt__.
  
### func
```
    func main(){

    }
```
- __func__ keyword for creating function .

### main.go file organised as!
 - package declaration .cd
 - import packages that we need.
 - declare functions