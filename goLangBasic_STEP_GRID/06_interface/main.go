package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englistBot struct{}
type spanishBot struct{}

func main() {
	eb := englistBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englistBot) getGreeting() string {
	return "Hie there!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}
