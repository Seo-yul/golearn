package main

import (
	"fmt"
)

func main() {
	str := "Hello World"
	// Hello World 문자코드
	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	fmt.Println(str)
	fmt.Println(string(runes))
}
