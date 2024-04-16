package publicpkg

import "fmt"

// Go 언어에서는 대소문자로 공개, 비공개가 정해짐
const (
	PI = 3.1415 // 공개되는 상수
	pi = 3.341516 
)

var ScreenSize int = 1000
var screenHeight int
func PublicFunc() {
	const MyConst = 100
	fmt.Println("This is public Function", MyConst)
}

func privateFunc() {
	fmt.Println("This is a private Function")
}

type MyInt int // 공개
type myString string

type MyStruct struct {
	Age int // 공개되는 필드
	name string
}

func (m MyStruct) PublicMethod() {
	fmt.Println("This is a public method")
}

func (m MyStruct) privateMethod() {
	fmt.Println("This is a private method")
}

type myPrivateStruct struct {
	Age int // 구조체가 private 여서 공개가 안되는 필드
	name string
}

func (m myPrivateStruct) PrivateMethod() {
	fmt.Println("This is a private method")
}
