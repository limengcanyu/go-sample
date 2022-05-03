package language

import "fmt"

func SliceNilMain() {
	fmt.Print("--------------------------------空(nil)切片--------------------------------\n")

	var numbers []int

	PrintSlice(numbers)

	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

func SliceCutMain() {
	fmt.Print("--------------------------------切片截取--------------------------------\n")

	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	PrintSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	PrintSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	PrintSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	PrintSlice(number3)

}

func SliceAppendMain() {
	fmt.Print("--------------------------------切片增加容量--------------------------------\n")

	var numbers []int
	PrintSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	PrintSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	PrintSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	PrintSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	PrintSlice(numbers1)
}

func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func SliceMain() {
	// 空(nil)切片
	SliceNilMain()

	// 切片截取
	SliceCutMain()

	// 增加切片容量
	SliceAppendMain()
}
