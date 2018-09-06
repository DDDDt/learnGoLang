package main

import (
	"fmt"
)

/*错误处理*/
type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {

	return fmt.Sprint("出错", de.dividee)

}

func Divide(varDividee int, varDivider int) (result int, errorMsh string) {

	if varDividee == 0 {
		var dDate = DivideError{varDividee, varDivider}
		return 0, dDate.Error()
	}

	return varDivider, ""

}

// 使用 panic recover defer
func f() {
	defer func() {
		fmt.Println("b")
		if err := recover(); err != nil {
			fmt.Println("c")
			fmt.Println(err)
		}
		fmt.Println("d")
	}()
	fmt.Println("a")
	panic("a bug occur")
	fmt.Println("e")

}

func main() {

	//errors.New("直接错误")
	//fmt.Println(Divide(0, 1))

	f()

}
