package main

import "fmt"

// 数组是值类型  go一般不直接使用数组

func printArray(arr [5]int) {
	arr[0] = 100
	for _, v := range arr { // 不想要某个值的话 就用下划线表示
		fmt.Println(v)
	}
}

func main() {
	var arr1 [5]int                  // 不用定义每一项
	arr2 := [3]int{1, 3, 5}          // 冒号定义的得写每一项得值
	arr3 := [...]int{2, 4, 6, 8, 10} // 不想数有几个
	var grid [4][5]int               // 4 个长度为 5的 int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}

	for i, v := range arr3 { // 同时获取下表和值
		fmt.Println(i, v)
	}

	for _, v := range arr2 { // 不想要某个值的话 就用下划线表示
		fmt.Println(v)
	}
	fmt.Println("printArray(arr1)")
	printArray(arr1)
	fmt.Println("printArray(arr3)")
	printArray(arr3)
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3) // 值并没有改变 array 作为参数的话 会拷贝数组
}
