package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#B71C1C",
	// 	"green": "#1B5E20",
	// 	"blue":  "#01579B",
	// }

	// var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["red"] = "#fdfdfd"
	colors["blue"] = "#fsjdflksjdf"

	// delete a map
	delete(colors, "red")

	fmt.Println(colors)
	fmt.Println(colors["red"])

	//print map
	printMap(colors)

}

//iterating over maps
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}

}
