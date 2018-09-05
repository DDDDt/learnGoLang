package main

import "fmt"

// 基本的运算符和其他语言相同

/*运算符*/
func main() {

	a := 1
	b := 2
	fmt.Println("相加 %d = ", a+b)

	// 返回变量存储地址
	println(&a)

	// 指针变量
	var ptr *int
	println(ptr)

}
