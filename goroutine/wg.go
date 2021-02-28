package main

import (
	"fmt"
	"sync"
)

//启10个goroutine
func wg() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int){
			if i < 5 {
				fmt.Println(i)
			}
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("finished ")

}
func main()  {
	wg()
}

// Tip: waitGroup 不要进行copy
func errWg1() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, wg sync.WaitGroup) {
			fmt.Println(i)
			defer wg.Done()
		}(i, wg)
	}
	wg.Wait()
	fmt.Println("finished")
}

// Tip: waitGroup 的 Add 要在goroutine前执行
func errWg2() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func(i int) {
			wg.Add(1)
			fmt.Println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("finished")
}

// Tip: waitGroup 的 Add 很大会有影响吗？
func errWg3() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(100)
		go func(i int) {
			fmt.Println(i)
			defer wg.Done()
			wg.Add(-100)
		}(i)
	}
	fmt.Println("finished")
}

//源码是一个for循环--过大影响性能

//for ; w != 0; w-- {
//	runtime_Semrelease(semap, false, 0)
//
//waiteGroup实现原理
//waiteGroup是两个信号量之间的交互