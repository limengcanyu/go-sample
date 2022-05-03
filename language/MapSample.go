package language

import "fmt"

func Map() {
	// 创建一个map，容量为10
	m := make(map[string]int, 1)

	// 添加元素
	m["001"] = 1
	m["001"] = 1
	m["002"] = 2
	m["003"] = 3

	for k, v := range m {
		fmt.Printf("element key: %s value: %d\n", k, v)
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	delete(m, "001")

	for k, v := range m {
		fmt.Printf("element key: %s value: %d\n", k, v)
	}

	length := len(m)
	fmt.Printf("element size: %d\n", length)
}
