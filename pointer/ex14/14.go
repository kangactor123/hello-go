package main

import "fmt"

type Data struct {
	value int
	data [200]int
}

func ex14_1() {
	var a int = 300
	var p *int

	p = &a

	fmt.Printf("p 의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)

	*p = 100
	fmt.Printf("a의 값: %d\n", a)
}

func ex14_2() {
	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	var p4 *int

	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p2 == p3 : %v\n", p2 == p3)

	// 포인터의 기본값은 nil 입니다.
	fmt.Println(p4, p4 == nil)
}

// call by address 인 경우 포인터를 받아야한다.
func changeData(args *Data) {
	args.data[100] = 100
	args.value = 100
}

func ex14_3() {
	data := Data{}
	fmt.Println(data.value)

	changeData(&data)
	fmt.Println(data.value)

	var p *Data = &Data{}
	var p2 = new(Data)

	p3 := &Data{} // p3은 포인터 변수임
}

func main() {
	ex14_1()
	ex14_2()
	ex14_3()

}