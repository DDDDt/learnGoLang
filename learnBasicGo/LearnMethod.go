package main

import "fmt"

// 方法

/*定义函数*/
type Circle struct {
	radius float64
}

// 定义属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {

	var c Circle
	c.radius = 10
	fmt.Println(c.getArea())

}
