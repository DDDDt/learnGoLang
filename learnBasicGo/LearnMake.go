package main

import "fmt"

/*切片*/

func main() {

	// [] 表示式切片类型
	s := []int{1, 2, 3}
	fmt.Println(s)

	// 将现有的数组转换为 切片
	var array = [4]int{1, 2, 3, 4}
	s1 := array[:]
	fmt.Println(s1)
	// 截取现有的数组转换为切片
	s2 := array[1:3]
	fmt.Println(s2)
	s3 := array[1:]
	s4 := array[:2]
	fmt.Println(s3)
	fmt.Println(s4)

	// 使用 make 初始化切片, 刚开始的切片默认是 0
	s5 := make([]int, 2, 3)
	fmt.Println(s5)
	fmt.Println(len(s5))
	fmt.Println(cap(s5))
	fmt.Println(s5 == nil)

	// 未定义默认的切片
	var numbers []int
	fmt.Println(numbers)
	fmt.Println(numbers == nil)

	// 切片截取
	newNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(newNumbers)
	fmt.Println(newNumbers[0])
	fmt.Println(newNumbers[3:])
	fmt.Println(newNumbers[:6])

	// 增加切片的容量
	var oldNumbers []int
	oldNumbers = append(oldNumbers, 1, 2, 3, 4)
	fmt.Println(oldNumbers)
	oldNumbers1 := make([]int, len(oldNumbers), cap(oldNumbers))

	fmt.Println(oldNumbers1)

	// 赋值切片
	copy(oldNumbers1, oldNumbers)
	fmt.Println(oldNumbers1)
	fmt.Println(oldNumbers)

}
