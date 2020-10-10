package main

import "fmt"

//Counter ...
type Counter int

func main() {
	fmt.Println("Write Counter")

	var c Counter
	var d0 []byte
	var d1 []byte
	var d2 []byte

	fmt.Printf("%+v\n", c)
	c.Write(d0)
	c.Write(d1)
	c.Write(d2)
	c.Write(d0)
	c.Write(d1)
	c.Write(d2)
	fmt.Printf("Writer Counter : %+v\n", c)

}
func (c *Counter) Write(p []byte) (n int, err error) {
	*c = Counter(int(*c) + 1)
	return len(p), nil
}
