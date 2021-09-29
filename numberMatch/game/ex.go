package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}
func main() {
	fmt.Printf("최댓값을 입력하세요:")
	n, err := InputIntValue()
	if err != nil {
		fmt.Printf("숫자만 입력가능")
	} else {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 6; i++ {
			v := rand.Intn(n)
			fmt.Printf("v: %v\n", v+1)
		}
	}

}
