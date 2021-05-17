package main

import (
	"go-sample/mylib"
)

func main() {
	// 空(nil)切片
	//mylib.SliceMain()

	// Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
	// 在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对的 key 值。
	//mylib.RangeMain()

	// 阶乘
	//mylib.FactorialMain()

	// 阶乘
	//mylib.FibonacciMain()

	// 接口
	//mylib.InterfaceMain()

	// 错误处理
	//mylib.ErrorMain()
	//
	//mylib.GoroutineMain()
	//
	//mylib.ChannelMain()
	//
	//mylib.ChannelBufferMain()
	//
	//mylib.ChannelRangeMain()

	// pulsar =============================================
	//pulsar.Producer()

	//pulsar.Consumer()

	// 使用监听器消费消息
	//pulsar.ConsumerListener()

	// chan =============================================
	// 方式1
	//go mylib.ChanSample()

	// 方式2
	naturals := make(chan int)
	squares := make(chan int)
	go mylib.Counter(naturals)
	go mylib.Squarer(squares, naturals)
	mylib.Printer(squares)

	// 当前线程休眠
	//time.Sleep(time.Second * 5)

	//mylib.Defer()
}
