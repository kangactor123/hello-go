package main

import "fmt"

// go 는 최 강타입 언어입니다. 타입이 다른 변수는 연산이 안됩니다.
func main() {
	fmt.Println("02.var")

	// a := 1

	// b := 2.3

	// c := a * b 정수와 실수는 연산할 수 없음

	// c := int(b)

	// d := a * c
	// fmt.Println(d)

	// 예제
	// e := 3
	var f float64 = 3.5

	// var g int = int(f)
	// h := float64(a * c)

	// var i int64 = 7
	// j := int64(d) * i
	var g1 int = int(f * 3)
	var h1 int = int(f) * 3
	fmt.Println(g1, h1);
}