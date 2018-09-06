package main

import "fmt"

func main() {

	// 使用randge 循环数组
	nums := [3]int{1, 2, 3}
	// 不使用索引时
	for _, n := range nums {
		fmt.Println(n)
	}
	for i, n := range nums {
		fmt.Println(i, n)
	}

	// 使用 range 循环 map
	kvs := map[string]string{"a": "apple", "b": "bb"}
	for k, v := range kvs {
		fmt.Println(k, v)
	}
}
