package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   // 4 byte
	Socre float64 // 8 byte
}

func main() {
	user := User{25, 99.9}
	fmt.Println(unsafe.Sizeof(user)) // 16
}
