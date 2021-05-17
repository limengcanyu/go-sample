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

func Counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func Squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func Printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
