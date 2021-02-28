package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main(){
	chClose()
}

func ch() {
	var ch = make(chan int)

	go func(ch chan int) {
		// Tip: 由于channel没有设置长度，所以是阻塞的，逐个发送
		ch <- 1
		ch <- 2
		ch <- 3
		fmt.Println("send finished")
	}(ch)

	for {
		select {
		case i := <-ch:
			fmt.Println("receive", i)
		case <-time.After(time.Second):
			fmt.Println("time out")
			os.Exit(1)
		}
	}
}

func chLimit() {
	var ch = make(chan int)

	// Tip: channel参数设置为 chan<- 和 <-chan，可以有效地防止误用发送和接收，例如这里的chan<-只能用于发送
	go func(ch chan<- int) {
		ch <- 1
		ch <- 2
		ch <- 3
		fmt.Println("send finished")
	}(ch)

	for {
		select {
		case i := <-ch:
			fmt.Println("receive", i)
		case <-time.After(time.Second):
			fmt.Println("time out")
			os.Exit(1)
		}
	}
}

func chClose() {
	var ch = make(chan int)

	go func(ch chan<- int) {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
		fmt.Println("send finished")
	}(ch)

	for {
		select {
		case i, ok := <-ch:
			if ok {
				fmt.Println("receive", i)
			} else {
				fmt.Println("channel close")
				os.Exit(0)
			}
		case <-time.After(time.Second):
			fmt.Println("time out")
			os.Exit(1)
		}
	}
}

func chCloseErr() {
	var ch = make(chan int)

	go func(ch chan<- int) {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
		fmt.Println("send finished")
	}(ch)

	for {
		select {
		// Tip: 如果这里不判断，那么i就会一直得到chan类型的默认值，如int为0，永远不会停止
		case i := <-ch:
			fmt.Println("receive", i)
		case <-time.After(time.Second):
			fmt.Println("time out")
			os.Exit(1)
		}
	}
}

func chTask() {
	var doneCh = make(chan struct{})
	var errCh = make(chan error)

	go func(doneCh chan<- struct{}, errCh chan<- error) {
		if time.Now().Unix()%2 == 0 {
			doneCh <- struct{}{}
		} else {
			errCh <- errors.New("unix time is an odd")
		}
	}(doneCh, errCh)

	select {
	// Tip: 这是一个常见的Goroutine处理模式，在这里监听channel结果和错误
	case <-doneCh:
		fmt.Println("done")
	case err := <-errCh:
		fmt.Println("get an error:", err)
	case <-time.After(time.Second):
		fmt.Println("time out")
	}
}

func chBuffer() {
	var ch = make(chan int, 3)

	go func(ch chan int) {
		// Tip: 由于设置了长度，相当于一个消息队列，这里并不会阻塞
		ch <- 1
		ch <- 2
		ch <- 3
		fmt.Println("send finished")
	}(ch)

	for {
		select {
		case i := <-ch:
			fmt.Println("receive", i)
		case <-time.After(time.Second):
			fmt.Println("time out")
			os.Exit(1)
		}
	}
}

func chBufferRange() {
	var ch = make(chan int, 3)

	go func(ch chan int) {
		// Tip: 由于设置了长度，相当于一个消息队列，这里并不会阻塞
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
		fmt.Println("send finished")
	}(ch)

	for i := range ch {
		fmt.Println("receive", i)
	}
}


// 示例1
type Ball struct {
	hits int
}

func passBall() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	// Tip: 核心逻辑：往channel里放入数据，作为启动信号；从channel读出数据，作为关闭信号
	table <- new(Ball)
	time.Sleep(time.Second)
	<-table
}

func player(name string, table chan *Ball) {
	for {
		// Tip: 刚进goroutine时，先阻塞在这里
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		// Tip: 运行到这里时，另一个goroutine在收数据，所以能准确送达
		table <- ball
	}
}

// 示例2
func passBallWithClose() {
	// Tip 虽然可以通过GC自动回收channel资源，但我们仍应该注意这点
	table := make(chan *Ball)
	go playerWithClose("ping", table)
	go playerWithClose("pong", table)

	table <- new(Ball)
	time.Sleep(time.Second)
	<-table
	close(table)
}

func playerWithClose(name string, table chan *Ball) {
	for {
		ball, ok := <-table
		if !ok {
			break
		}
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

// 示例3
type sub struct {
	// Tip 把chan error看作一个整体，作为关闭的通道
	closing chan chan error
	updates chan string
}

func (s *sub) Close() error {
	// Tip 核心逻辑：两层通知，第一层作为准备关闭的通知，第二层作为关闭结果的返回
	errc := make(chan error)
	// Tip 第一步：要关闭时，先传一个chan error过去，通知要关闭了
	s.closing <- errc
	// Tip 第三步：从chan error中读取错误，阻塞等待
	return <-errc
}

func (s *sub) loop() {
	var err error
	for {
		select {
		case errc := <-s.closing:
			// Tip 第二步：收到关闭后，进行处理，处理后把error传回去
			errc <- err
			close(s.updates)
			return
		}
	}
}


//底层实现 ch的接受和阻塞是注册在一个结构体中的
