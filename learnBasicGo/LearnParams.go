package main

/*变量赋值*/

/*指定变量类型*/
var name string = "dt"
var nameDefault string

/*默认是0*/
var ageDefault int

/*自动类型*/
var v_name = "123"

/*省略 var , 报错, 只能出现在函数体内，不可以用于全局变量的声明和赋值 */
//c := 10

/*多变量声明*/
var vmname1, vmname2, vmname3 string

var vmname4, vmname5, vmname6 = 1, 2, "dt"

/*这种因式分解关键字的写法一般用于声明全局变量*/
var (
	a int
	b bool
)

func main() {
	c := 10
	d, e, f := 1, 2, "dd"

	/*进行值交换*/
	d, e = e, d
	println(name, nameDefault, ageDefault, v_name, c)
	println(vmname1, vmname2, vmname3, vmname4, vmname5, vmname6, d, e, f)
	println(a, b)
	a = 1
	b = false
	println(a, b)
}
