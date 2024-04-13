package main

import "fmt"

func main() {
	// 키워드 변수명 타입 = 값
	var a int = 10
	var msg string = "hello go world"

	fmt.Println(a)

	a = 30

	fmt.Println(a)

	fmt.Println(msg, a);

	var minWage int = 10
	var workingHour int = 20

	var income int = minWage * workingHour

	fmt.Println(minWage, workingHour, income)

	var i = 10 // 타입추론

	fmt.Println(i);

	// 선언대입문(:=) 를 활용해 변수 키워드, 타입을 생략
	k := "hi"

	fmt.Println(k);

	// const CONST int = 10

	const (
		Red int = iota
		Blue int = iota	
		Green int = iota
	)

	fmt.Println(Red, Blue, Green)


}