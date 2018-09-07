package main

import (
	"fmt"
)

/*channel*/

func main() {

	// chan 里面放 interface 类型
	var userChan chan interface{}
	userChan = make(chan interface{}, 10)

	userChan <- "nick"
	userChan <- "123"
	userChan <- "1213"
	userChan <- "1223"
	userChan <- "1233"
	userChan <- "1243"
	userChan <- "1253"
	userChan <- "1263"
	userChan <- "1273"
	userChan <- "1283"

	name, ok := <-userChan

	fmt.Println(name, ok)

	// 关闭渠道,关闭后在放入数据会 panic: send on closed channel . chan 不关闭取超数据的情况会报 deadlock
	close(userChan)
	select {
	case m := <-userChan:
		fmt.Println(m)
	default:
		fmt.Println("........")
	}
	// 要放在 close 之后才可以
	for v := range userChan {
		fmt.Println(v)
	}

	/*	weight := 55
		height := 1.70*1.70
		bmi := float64(weight)/height
		go fmt.Println(bmi)
		switch {
		case bmi < 18.5:
			fmt.Println("偏轻")
		case 18.5 <= bmi,bmi <= 24:
			fmt.Println("体重正常")
		case 25 <= bmi,bmi <= 30:
			fmt.Println("体重正常")
		case 30 <= bmi,bmi <= 35:
			fmt.Println("体重正常")
		default:
			fmt.Println(bmi)
		}

		time.Sleep(time.Second)
	*/
}
