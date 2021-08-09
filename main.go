package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for { // 무한 루프
		fmt.Println("input")
		var number int
		_, err := fmt.Scanln(&number) // 한 줄 입력
		if err != nil {               // 숫자가 아니면
			fmt.Println("input number")

			stdin.ReadString('\n') // 키보드 버퍼 제거
			continue               // for문 처음으로
		}
		fmt.Printf("input number is %d \n", number)
		if number%2 == 0 { // 짝수 검사
			break // for문 종료
		}
	}
	fmt.Println("for statement is end")
}
