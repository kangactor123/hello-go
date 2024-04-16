package initpkg

import "fmt"

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

// init 함수는 패키지를 임포트 할 때 실행을 해 패키지를 초기화 합니다.
func init() {
	d++
	fmt.Println("init func", d)
}

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}

func PrintD() {
	fmt.Println("d:", d)
}