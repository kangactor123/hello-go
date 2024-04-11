package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

// 멀티 반환 함수
func multiAdd(a int, b int) (int, bool) {
	return a + b, a > b
}

// 이름을 지정한 멀티 반환 함수
func divide(a int, b int) (result int, success bool) {
	if (b == 0) {
		result = 0
		success = false
		return
	}

	result = a / b
	success = true
	return
}

func selfFunc(a int, b int) {
	if (a == 0 || b == 0) {
		fmt.Print("a || b cant zero")
		return
	}

	if (a % b == 0) {
		fmt.Println("Finally 0")
		return
	} else {
		fmt.Println(a / b)
		selfFunc(a, a % b)
	}
}

func main() {
	var a int
	var b int
	n, err := fmt.Scan(&a, &b)

	if (err != nil) {
		fmt.Println(n, err)
	} else {
		// sum := add(a, b);
		// fmt.Println("n: ", n, ", sum: ", sum)
		selfFunc(a, b)
	}

	// c, d := multiAdd(a, b)

	// fmt.Println("multiAdd: ", c, d)

	// e, f := divide(a, b)

	// fmt.Println("divide: ", e, f)
}