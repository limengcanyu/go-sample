package language

import "fmt"

func SwitchSample() {
	// Expression switches
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

	// 一分支多值
	var b = "mum"
	switch b {
	case "mum", "daddy":
		fmt.Println("family")
	}

	// 分支表达式
	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}

	// Type switches
	//var x interface{} = "A sample string" // 空接口类型（interface{}）
	//var x interface{} = 0.1
	var x interface{} = nil
	fmt.Println(t(x)) // A string value
}

func t(i interface{}) string {
	switch i.(type) {
	case string:
		return "A string value"
	case int:
		return "A number"
	case nil:
		return "A nil"
	default:
		return "Other type"
	}
}
