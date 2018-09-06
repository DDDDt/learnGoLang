package main

import "fmt"

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

func main() {

	fmt.Println(Divide(0, 1))

}
