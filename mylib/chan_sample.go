package mylib

import "fmt"

func ChanSample() {
	// 构建通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start send goroutine")

		// 向通道发送值
		ch <- 0

		fmt.Println("exit send goroutine")
	}()

	fmt.Println("wait receive goroutine")

	// 从通道接收值
	value := <-ch
	fmt.Printf("receive int value: %d\n", value)

	fmt.Println("all done")
}

func Send(ch chan int) {
	fmt.Println("start send goroutine")

	// 向通道发送值
	ch <- 1

	fmt.Println("exit send goroutine")
}

func Receive(ch chan int) {
	fmt.Println("start receive goroutine")

	// 从通道接收值
	value := <-ch
	fmt.Printf("receive int value: %d\n", value)

	fmt.Println("exit receive goroutine")
}
