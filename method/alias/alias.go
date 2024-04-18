package alias

type MyInt int

func (a MyInt) Add(b int) int {
	return int(a) + b
}