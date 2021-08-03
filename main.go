package main

import "fmt"

func main() {
	day := "thursday"

	switch day {
	case "monday", "tuesday":
		fmt.Println("월, 화")
	case "wedneday", "thurday", "friday":
		fmt.Println("수, 목, 금")
	}
}
