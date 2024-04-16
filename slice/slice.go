package main

import (
	"fmt"
)

func slicing() {
	array := [5]int{1, 3, 5, 6, 8}

	// 슬라이싱 [startIdx:endIdx - 1]
	slice := array[1:5]

	fmt.Println(array)
	fmt.Println(slice)

	// slicing 해도 해당 주소를 바라보고 있음
	slice[0] = 100

	fmt.Println(array)
	fmt.Println(slice)

	slice2 := array[:4]
	slice3 := array[3:]

	fmt.Println(slice2)
	fmt.Println(slice3)
}

// 슬라이스 복제
func duplicateSlice() {
	slice := []int{1, 3, 5, 7}

	slice2 := make([]int, len(slice))

	fmt.Println(slice2)
	// for i, v := range slice {
	// 	slice2[i] = v
	// }

	slice3 := append([]int{}, slice...) // 깊은복사

	copy(slice2, slice) // slice2 에 slice 를 복사

	slice[1] = 100
	fmt.Println(slice)
	fmt.Println(slice3)
	fmt.Println(slice2)
}

func addElement() {
	idx := 2
	slice := []int{1,2,3,4}
	slice = append(slice, 0)

	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}

	slice[idx] = 100

	slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)
	fmt.Println(slice)
}

func main() {
	fmt.Println("")

	var slice []int

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	// slice[1] = 10 // 패닉 발생
	// fmt.Println(slice)

	// array := [...]int{1,3,5} // 길이가 3인 배열
	// slice2 := []int{1,3,5} // 현재 길이가 3인 슬라이스

	// var slice3 = make([]int, 3) // make 함수를 활용해 slice 초기화

	// for idx, n := range slice3 {
	// 	fmt.Println(n)
	// 	slice3[idx] = n + 3
	// 	fmt.Println(slice3[idx])
	// }

	// var slice4 = []int{1, 3, 5}

	// slice4 = append(slice4, 4)

	// fmt.Println(slice4)

	// test := []int{}

	// for i := 0; i < 10; i++ {
	// 	test = append(test, i)
	// }

	// fmt.Println(test)

	// fmt.Println("slicing\n")

	// slicing()

	duplicateSlice()

 }