package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"miss", "kiss", "tiss"}
	fmt.Println(mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
