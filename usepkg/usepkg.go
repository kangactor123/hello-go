package main

// 패키지를 임포트 한다면 생기는 일
// 1. 패키지의 전역 변수를 초기화한다.
// 2. init 함수를 호출해 패키지를 초기화 한다. (init 함수는 리턴과 매개변수가 없는 함수여야한다)
// 초기화 기능만 이용하기 원할 경우 패키지를 import 할 때 _ 를 사용한다.
import (
	"fmt"
	"goproject/usepkg/custompkg"
	"goproject/usepkg/initpkg"
	"goproject/usepkg/publicpkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3,4,5,6,8,1,2,3,4,1,2}
	graph := asciigraph.Plot(data)

	fmt.Println(graph)

	fmt.Println("PI:", publicpkg.PI)
	publicpkg.PublicFunc()

	var myInt publicpkg.MyInt = 10
	fmt.Println("myint:", myInt)

	var MyStruct = publicpkg.MyStruct{Age: 18}
	fmt.Println("MyStruct:", MyStruct)

	initpkg.PrintD()
}