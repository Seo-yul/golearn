package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 1 byte
	C int8
	B int // 8 byte
	D int
}

func main() {
	user := User{1, 2, 3, 4}
	fmt.Println(unsafe.Sizeof(user)) // 24
}
