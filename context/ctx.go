package main

import (
	"context"
	"fmt"
	"time"
)

// Tip: 通过 cancel 主动关闭
func ctxCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(time.Millisecond * 100):
			fmt.Println("Time out")
		}
	}(ctx)

	cancel()
}

// Tip: 通过超时，自动触发
func ctxTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	// 主动执行cancel，也会让协程收到消息
	defer cancel()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(time.Millisecond * 100):
			fmt.Println("Time out")
		}
	}(ctx)

	time.Sleep(time.Second)
}

// Tip: 通过设置截止时间，触发time out
func ctxDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Millisecond))
	defer cancel()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(time.Millisecond * 100):
			fmt.Println("Time out")
		}
	}(ctx)

	time.Sleep(time.Second)
}

// Tip: 用Key/Value传递参数，可以浅浅封装一层，转化为自己想要的结构体
func ctxValue() {
	ctx := context.WithValue(context.Background(), "user", "qlf")
	go func(ctx context.Context) {
		v, ok := ctx.Value("user").(string)
		if ok {
			fmt.Println("pass user value", v)
		}
	}(ctx)
	time.Sleep(time.Second)
}

func main() {
	ctxValue()
}


//把当前停止的进程的一些信息存储下来传递到下一个进程    类比cpu上下文切换和子进程和父进程的关系（消息通知）
//优雅的实现 select
//Context的底层实现 mutex和channel的结合 前者用于初始部分参数，和者用于通信   (go语言编程内涵，不要用内存通信,使用channel通信）