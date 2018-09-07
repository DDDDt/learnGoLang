package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int
}

/*反射*/
func main() {

	a := 1

	// 检查是否可以被修改
	vof := reflect.ValueOf(a)
	// int
	fmt.Println(vof.Kind().String())
	// 1
	fmt.Println(vof.Int())
	// 1 原接口的值
	fmt.Println(vof.Interface())

	// false
	fmt.Println(vof.CanSet())
	// 报错 unaddressable value, 需要检查是否可以被修改
	//reflect.ValueOf(a).SetLen(10)

	// 想要改变,就必须传递指针类型
	ptrOf := reflect.ValueOf(&a)
	fmt.Println(ptrOf.CanSet())
	fmt.Println(ptrOf.Elem())
	fmt.Println(ptrOf.Elem().CanSet())
	ptrOf.Elem().SetInt(10)
	fmt.Println(ptrOf.Interface())
	fmt.Println(a)
	// 得到 type 类型
	of := reflect.TypeOf(a)
	fmt.Println(of.Name())

	/*反射后的类型是一个结构类型*/

	// 定义 object 类型

}
