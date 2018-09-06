package main

import "fmt"

/*map*/

func main() {

	var kvs map[string]string
	var kvm = make(map[string]string)

	fmt.Println(kvs)
	fmt.Println(kvm)

	/*nil map 不能用来存放键值对,会报错*/
	//kvs["a"] = "apple"

	kvm["a"] = "apple"
	fmt.Println(kvs)
	fmt.Println(kvm)

	for k, v := range kvm {
		fmt.Println(k, v)
	}

	/*查看元素在集合中是否存在*/
	v, ok := kvm["a"]
	_, ok1 := kvm["b"]
	fmt.Println(v, ok)
	fmt.Println(ok1)

	// 删除 delete
	delete(kvm, "a")

	// 删除 nil 也不会报错
	delete(kvs, "a")

	fmt.Println(kvm)

}
