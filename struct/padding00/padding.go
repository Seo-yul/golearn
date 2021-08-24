package main

import "fmt"

type Student struct {
	Age    int // public
	Number int
	Score  float64
}

func PrintStudent(s Student) {
	fmt.Printf("age: %d number: %d, score: %.f\n", s.Age, s.Number, s.Score)
}

func main() {
	var s = Student{25, 5, 99.9}

	s2 := s // copy s value

	s.Age = 10

	PrintStudent(s2)
	// age: 25 number: 5, score: 100
}
