package main

import "fmt"

func main() {
	// 큰따옴표로 묶으면 특수 문자가 기능을 한다.
	str1 := "Hello\t'world'\n"

	// 백쿼트로 묶으면 특수 문자가 동작하지 않는다.
	str2 := `Go is "awesome" \nGo is simple and\t 'powerful'`

	fmt.Print(str1)
	fmt.Print(str2)
}
