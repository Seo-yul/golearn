package main

import (
	"fmt"
)

func getMyAge() int {
	return 11
}

func main() {
	// age 변수 선언 및 초기화
	switch age := getMyAge(); { // 비굣값이 true이기 때문에 생략
	case age < 10:
		fmt.Println("child")
	case age < 20:
		fmt.Println("teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)
	}
}
