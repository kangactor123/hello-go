package main

import "fmt"

const ARRAY_LENGTH = 5

func ex12_1() {
 	// 배열 선언
	// [길이]타입 = [길이]타입{}
	var array [5]int = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	// 배열의 길이는 항상 상수로 넣어야 함
	var array2 [ARRAY_LENGTH]int = [ARRAY_LENGTH]int{1,3,5,6,7}

	for i := 0; i < len(array2); i++ {
		if (array2[i] % 2 == 0) {
			fmt.Println(array2[i], i, "짝수입니다.")
		}
	}
}

func ex12_2() {
	nums := [...]int{1,3,5,7,9}

	nums[2] = 300

	for i := 0; i < len(nums); i++ {
		if nums[i] > 100 {
			fmt.Println(nums[i], "은 100보다 큽니다.")
			break
		}
	}
}

// 배열 순회
func ex12_3() {
	var array = [...]int{1, 3, 5}

	// range 키워드를 활용해 배열을 순회함
	// i 는 index, v 는 해당 인덱스의 값
	for i, v := range array {
		fmt.Println(i, v)
	}
}

func ex12_4() {
	a := [ARRAY_LENGTH]int{1,2,3,4,5}
	b := [ARRAY_LENGTH]int{6,7,8,9,10}

	for _, v := range a {
		fmt.Println("a", v)
	}

	fmt.Println()

	for _, v := range b {
		fmt.Println("b", v)
	}
	
	fmt.Println()

	// 배열의 대입은 전체 배열의 복사로 동작
	b = a
	
	for _, v := range b {
		fmt.Println("b", v)
	}
}

func main() {
	// ex12_1()
	// ex12_2()
	// ex12_3()
	ex12_4()
}