package main

import "unsafe"

/*定义常量关键字, 类似于 scala 的 val*/

const cname string = "dt"

const cname1 = "dt1"

/*常量必须要有初始值*/
//const cname1 string

const c_name, c_name1, c_name2 = 1, 2, "dt1"

// 使用 const 定义枚举
const (
	Name = "dt"
	Age  = 25
)

// 使用 iota
const (
	i = iota
	i1
	i2
	i3
)

const (
	N = iota
	N1
	N2
)

func main() {
	println(cname, cname1)
	println(c_name, c_name1, c_name2)
	println(Name, Age)
	println(unsafe.Sizeof(Name))

	println(i, i1, i2, i3)

	println(N, N1, N2)

}
