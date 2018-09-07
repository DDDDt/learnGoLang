package main

import (
	"encoding/json"
	"fmt"
)

/*json*/

// omitempty 过滤某些元素
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func main() {

	per := person{"dt", 25}
	fmt.Println(per)

	// 转换为 json, 注意 结构体的属性要大写
	valuer, err := json.Marshal(per)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(valuer))

	var pers person

	// json 转换为 struct
	json.Unmarshal(valuer, &pers)

	fmt.Println(pers)

}
