package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func getRandomIntn(rangeInt int) int {
	random := rand.Intn(rangeInt)

	fmt.Println("랜덤값:", random)
	return random
}

func InputIntValue() (int, error) {
	var input int
	
	_, err := fmt.Scanln(&input)

	if err != nil {
		stdin.ReadString('\n') //애러 발생 시 입력 스트림을 비움
	}

	return input, err
}

func main() {
	fmt.Println("숫자 맞추기 게임에 오신걸 환영합니다!")

	flag := true
	answer := getRandomIntn(100)

	for flag {
		var input int
		
		fmt.Print("숫자를 입력해주세요: ")
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("에러가 존재합니다!")
			continue
		} else {
			fmt.Println("입력하신 숫자는 다음과 같습니다", input)

			if answer > input {
				fmt.Println("입력하신 숫자가 정답보다 작습니다.")
			} else if answer < input {
				fmt.Println("입력하신 숫자가 정답보다 큽니다.")
			} else {
				fmt.Println("정답입니다!!", answer)
				break;
			}				
		}
	}


}