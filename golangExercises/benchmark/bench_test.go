package main

import (
	"fmt"
	"testing"
)

func BenchTest(b *testing.B) {
	for n := 0; n < 1000; n++ {
		fmt.Println(n)
	}
}
