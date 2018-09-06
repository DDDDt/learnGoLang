package main

import (
	"fmt"
	"reflect"
)

/*类型转换*/

func main() {

	sum := 10
	fmt.Println(reflect.TypeOf(sum))

	var f float32 = float32(sum)

	fmt.Println(reflect.TypeOf(f))

}
