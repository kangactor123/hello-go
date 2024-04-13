package main

import "fmt"

func main() {
	a := 1
	b := 1

//레이블 정의
Outerfor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if (a * b == 45) {
				// 레이블에 가장 먼저 포함된 for문 종료
				break Outerfor
			}
		}
	}
	
	fmt.Printf("%d * %d = %d\n", a, b, a * b)
}