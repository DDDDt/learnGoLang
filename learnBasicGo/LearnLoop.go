package main

import "fmt"

/*循环判断语句*/
func main() {

	// 无限循环
	/*for true {
		fmt.Println("这是无限循环操作")
	}*/

	/* for 循环*/
	for i := 0; i < 10; i++ {
		fmt.Println("i 的值 %d", i)
	}

	i := 0
	for i < 10 {
		fmt.Println("i 的值 %d", i)
		i++
	}

	// 这是一个死循环
	/*for{
		fmt.Println("一直进来")
	}*/

	/*嵌套循环*/
	for i := 0; i < 100; i++ {
		if i == 88 {
			continue
		}
		for n := 0; n < 100; n++ {
			fmt.Println("i 的值 %d,n的值 %d", i, n)
		}
		if i == 99 {
			break
		}
	}

}
