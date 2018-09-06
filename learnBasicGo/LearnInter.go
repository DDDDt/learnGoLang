package main

import "fmt"

/*接口*/
type inter interface {
	test()
	max(a, b int) (c int)
}

type test struct {
}

func (t test) test() {
	fmt.Println(123)
}

func (t test) max(a, b int) (c int) {

	if a > b {
		return a
	} else {
		return b
	}

}

func main() {

	t := new(test)
	t.test()
	fmt.Println(t.max(1, 2))

}
