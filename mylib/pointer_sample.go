package mylib

import "fmt"

func PointerSample() {
	var cat int = 1
	var str string = "banana"
	fmt.Printf("address cat: %p str: %p\n", &cat, &str)

	var i *int = nil // 定义int类型指针
	//fmt.Printf("i: %d\n", *i) // cause a run-time panic
	j := 1 // 定义int类型变量
	i = &j // 将j的内存地址赋值给i
	fmt.Printf("i: %d\n", i)
	fmt.Printf("i: %d\n", &i)
	fmt.Printf("i: %d j: %d\n", *i, j)

	*i = 2
	fmt.Printf("i: %d j: %d\n", *i, j)

	// 创建指针的另一种方法——new() 函数
	str2 := new(string)
	*str2 = "Go语言教程"
	fmt.Println(*str2)
}
