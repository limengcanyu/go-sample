package language

import (
	"fmt"
)

func Defer() {
	//// prints 3 2 1 0 before surrounding function returns
	//for i := 0; i <= 3; i++ {
	//	defer fmt.Println(i)
	//}
	//
	//result := f()
	//fmt.Printf("result: %d\n", result)

	// 按照定义的相反顺序依次执行
	defer f1()
	defer f2()
	defer f3()
}

// f returns 42
func f() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}

func f1() (result int) {
	fmt.Println("call f1")
	return 1
}

func f2() (result int) {
	fmt.Println("call f2")
	return 2
}

func f3() (result int) {
	defer f4()
	defer f5()

	fmt.Println("call f3")
	var res int
	return res
}

func f4() (result int) {
	fmt.Println("call f4")
	return 4
}

func f5() (result int) {
	fmt.Println("call f5")
	return 5
}
