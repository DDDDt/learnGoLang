package main

import "fmt"

/*递归*/

func Factorial(x int) (result int) {

	if x == 0 {
		return x
	}
	x--
	Factorial(x)
	return
}

func main() {

	fmt.Println(Factorial(10))

}
