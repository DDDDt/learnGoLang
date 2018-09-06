package main

import "fmt"

/*数组*/

// 向函数中传递数组,固定长度
func arrays(a [10]int) {
	fmt.Println(a)
}

// 不固定长度
func arrayss(a []int) {
	fmt.Println(a)
}

func main() {

	// 定义数组
	var nums = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// 不设置长度，根据元素的个数来设置数组的大小
	var numss = []float32{1000.0, 2.0, 3.4, 7.0, 1}

	fmt.Println(nums)
	fmt.Println(numss)
	println(len(nums))

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// 多为数组

	a := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11}}

	fmt.Println(a)

	// 向函数中传递数组

}
