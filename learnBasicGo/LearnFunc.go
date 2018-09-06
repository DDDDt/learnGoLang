package main

import (
	"fmt"
	"math"
)

/*go 函数*/

//返回两个数的最大值
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 函数返回多个值, 类似于其他函数的元组
func swap(x, y string) (string, string) {

	return y, x

}

// 值传递，go 函数默认是值传递, 不会影响本来的数据
func swapValue(x, y string) (string, string) {
	var tmp = x
	x = y
	y = tmp
	return x, y
}

// 使用引用传递,会改变本来的值
func swapPoint(x, y *string) {
	temp := *x
	*x = *y
	*y = temp
}

// go 函数闭包
func getSequences() func() int {

	i := 0
	return func() int {
		i += 1
		return i
	}

}

func main() {

	num1 := 1
	num2 := 2
	fmt.Println(max(num1, num2))

	x := "1"
	y := "2"
	var a, b = swap(x, y)
	fmt.Println(a, b)

	c, d := swapValue(x, y)
	fmt.Println(x, y)
	fmt.Println(c, d)

	// 得到调用指针
	fmt.Println(&x, &y)
	// 得到调用指针的值
	fmt.Println(*&x, *&y)

	swapPoint(&x, &y)
	fmt.Println(x, y)

	// 将函数作为值, 类似于匿名函数
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))

	// 闭包, 类似于 js 的闭包
	fmt.Println(getSequences()())
	fmt.Println(getSequences()())
	fmt.Println(getSequences()())

}
