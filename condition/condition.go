package main

import "fmt"

func isAdult() (int, bool) {
	var age int
 	_, err := fmt.Scanln(&age)

	// 초기식이 있는 조건문
	if success := false; err != nil {
		success = false
		return 0, success
	} else {
		success = true
		return age, age > 18
	}
}

type ColorType int
const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func getColorType(color ColorType) {
	switch color {
	case Red:
		fmt.Println("Red")
	case Blue:
		fmt.Println("Blue")
	case Green:
		fmt.Println("Green")
	case Yellow:
		fmt.Println("Yellow")
	default:
		fmt.Println("Default")
	}
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}
}

func main() {
	age, isAdult := isAdult();

	if (age == 0) {
		fmt.Println("Err")
	} else {
		fmt.Println(isAdult, age)
	}

	a := 3

	switch a {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println("default")
	}

	getColorType(1)

	loop()
}