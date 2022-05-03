package language

import "fmt"

func AssignmentsSample() {
	// 定义函数，返回2个结果
	f := func() (a int, b int) {
		return 1, 2
	}

	x, y := f() // 将函数结果赋值给x,y

	fmt.Printf("x: %d y: %d\n", x, y)
}
