package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	n, err := fmt.Println("hello")
	if err != nil {
		fmt.Println(err)
	}
	m, err := fmt.Println("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n, m)
	///
	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err = fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}
	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}
	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer1, answer2, answer3)
	/// create file
	f, err := os.Create("names.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := strings.NewReader("James Bond")
	io.Copy(f, r)

	/// open file
	ff, err := os.Open("names.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer ff.Close()

	bs, err := ioutil.ReadAll(ff)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
