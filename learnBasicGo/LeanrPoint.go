package main

import "fmt"

/*指针*/

// 向函数传递指针
func ptrMethod(a, b *int) (int, int) {

	return *a, *b

}

func main() {

	var a int = 20
	// 声明指针变量
	var ip *int
	ip = &a
	println(ip)
	// 访问指针变量的值
	println(*ip)
	/*空指针*/
	var ptr *int = nil
	println(ptr)
	fmt.Println(ptr)

	// 指针数组
	var arr = [4]int{1, 2, 3, 4}
	var ptrArr [4]*int
	for i := 0; i < len(arr); i++ {
		ptrArr[i] = &arr[i]
	}

	fmt.Println(ptrArr)

	for i := 0; i < len(ptrArr); i++ {
		fmt.Println(*ptrArr[i])
	}

	// 指向指针的指针
	var ptrPtr **int
	tmpPtr := &a
	ptrPtr = &tmpPtr
	fmt.Println(ptrPtr)
	fmt.Println(**ptrPtr)

	// 向函数传递指针
	fmt.Println(ptrMethod(&a, &a))

}
