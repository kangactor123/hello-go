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

// return isAnswer
func isAnswer(input int, answer int) bool {
	flag := false

	if input == answer {
		fmt.Println("정답입니다!!", answer)
		flag = true
	} else if input > answer {
		fmt.Println("입력하신 숫자가 정답보다 큽니다.")
	} else {
		fmt.Println("입력하신 숫자가 정답보다 작습니다.")
	}

	return flag;
}

func main() {
	fmt.Println("숫자 맞추기 게임에 오신걸 환영합니다!")

	flag := true
	answer := getRandomIntn(100)

	for flag {
		fmt.Print("숫자를 입력해주세요: ")
		value, err := InputIntValue()

		if err != nil {
			fmt.Println("숫자만 입력하세요.")
			continue
		}

		isAnswer := isAnswer(value, answer)

		if isAnswer {
			fmt.Println("프로그램을 종료합니다.")
			break
		}
	}


}