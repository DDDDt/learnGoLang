package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*Goroutine*/

var (
	m = make(map[int]uint64)
	// 申明一个互斥锁
	lock sync.Mutex
)

func domainPut(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error to chan put.")
		}
	}()
	panic("error.....")
}

type task struct {
	n int
}

func cals(t *task) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error....")
		}
	}()

	var sum uint64
	sum = 1
	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}
	lock.Lock() // 写全局数据加互斥锁
	m[t.n] = sum
	lock.Unlock() // 解锁

}

func wcal(w *sync.WaitGroup, i int) {

	fmt.Println("calc : ", i)
	time.Sleep(time.Second)

	// 添加到sync
	w.Done()

}

func main() {

	// 本地机器的逻辑 CPU 个数
	num := runtime.NumCPU()
	fmt.Println(num)
	// 设置可同时执行的最大 CPU 数, 并返回先前的设置
	runtime.GOMAXPROCS(num)

	// 使用 go
	go domainPut(123)
	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		t := task{10}
		go cals(&t)
	}

	time.Sleep(time.Second * 10)

	lock.Lock()
	for k, v := range m {
		fmt.Println(k, v)
	}
	lock.Unlock()

	// 等待所有任务退出，主程序再退出
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go wcal(&wg, i)
	}
	wg.Wait()
	fmt.Println("结束了...")

}
