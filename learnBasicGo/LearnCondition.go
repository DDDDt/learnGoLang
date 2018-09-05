package main

/*条件语句*/
func main() {

	// if 判断语句
	a := 20
	if a <= 20 {
		println("a 小于 20")
	}

	// if else 语句
	if a < 20 {
		println("a 小于 20")
	} else {
		println("a 大于等于 20")
	}
	// if else 嵌套
	if a <= 20 {
		if a == 20 {
			println("a 等于 20")
		}
	}
	// switch case
	switch a {
	case 20:
		println("a 等于 20")
	case 21:
		println("a 等于 21")
	default:
		println("进入默认值")
	}
	// type switch
	var x interface{}
	switch x.(type) {
	case string, bool:
		println("x的字符串或者是bool")
	case int:
		println("x是int类型")
	default:
		println("默认值类型")
	}

	// select case 监听 IO 操作,会一直等待某个 case 语句完成才会退出
	select {}

}
