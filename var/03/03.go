package main

import "fmt"

var PACKAGE int = 3

func main() {
	fmt.Println(PACKAGE)
	
	{
		var local int = 4

		fmt.Println(local)
	}

	var a int
	var b int

	// n, err := fmt.Scan(&a, &b)

	// if (err != nil) {
	// 	fmt.Println(n, err)
	// } else {
	// 	fmt.Println(n, a, b)
	// }

	n, err := fmt.Scanln(&a, &b) // 공란으로 구분된 값을 읽는다.

	if (err != nil) {
		fmt.Println(n, err)
	} else {
		fmt.Println(a, b, n)
	}

	// scope 를 벗어남
	// fmt.Println(local) 
}