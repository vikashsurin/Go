package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://golang.org",
		"https://amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// fmt.Println(<-c) // outPut value from channel
	// fmt.Println(<-c) // outPut value from channel

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c) // string and channel as argument
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down.")
		c <- link // pass value in channel
		return
	}
	fmt.Println(link, " is up.")
	c <- link // pass value in channel
}
