package main

import (
	"fmt"
	"strings"
)

// 인덱스를 통해 바이트 단위 순회
func visitString_1() {
	str := "Hello 월드!"

	for i := 0; i < len(str); i++ {
		fmt.Printf(" 타입: %T 값: %d 문자값: %c\n", str[i], str[i], str[i])
	}
}

// rune 슬라이스로 변환해 순회
func visitString_2() {
	str := "Hello 월드!"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf(" 타입: %T 값: %d 문자값: %c\n", arr[i], arr[i], arr[i])
	}
}

// range 키워드를 사용해 한 글자씩 순회
func visitString_3() {
	str := "Hello 월드!"

	for _, value := range str {
		fmt.Printf(" 타입: %T 값: %d 문자값: %c\n", value, value, value)
	}
}

func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' +  (c - 'a'))
		} else {
			rst += string(c)
		}
	}

	return rst
}

// 문자열은 불변이다. string 합 연산을 빈번하게 사용한다면 메모리를 낭비할 수 있다.
// string.Builder 을 활용해 메모리 낭비를 줄일 수 있음.
func ToUpper2(str string) string {
	// strings.Builder는 내부에 슬라이스를 가지고 있기 때문에 
	// WriteRune() 메서드를 통해 문자를 더할 때 매번 메모리를 생성하지 않고
	// 기존 메모리 공간에 빈자리가 있으면 그냥 더하게 된다.
	var builder strings.Builder

	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}

	return builder.String()
}

func stringBuilder() {
	str := "Hello World!"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))

}

func main() {
	// 하나의 문자를 담는 타입 > rune
	// rune 타입은 int32
	var char rune = '한'

	fmt.Printf("%T\n", char)
	fmt.Println(char)
	fmt.Printf("%c\n", char)

	a := "가나다라마"
	b := "abcde"

	// len 함수에 문자열을 넣으면 길이가 아닌 차지하는 메모리 크기를 리턴한다.
	fmt.Println(len(a)) // 한글은 3byte
	fmt.Println(len(b)) // 영문과 숫자는 1byte

	str := "Hello World!"
	runes := []rune(str) // rune 슬라이스로 문자열을 담는다.

	fmt.Println(str, runes)
	fmt.Println("str len:", len(str))
	fmt.Println("runes len:", len(runes))
	fmt.Println("length:", len([]rune(str)))

	// visitString_1()
	// visitString_2()
	visitString_3()

	stringBuilder()

}