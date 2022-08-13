package main

import "fmt"

func printArray(arr *[3]int) { // 带数字的[]表示的是数组
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}   // 这种声明方式必须赋初始值
	arr3 := [...]int{2, 4, 6} // 就是让编译器帮我们数有多少个
	fmt.Println(arr1, arr2, arr3)

	var grid [4][5]int // 4行5列
	fmt.Println(grid)

	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 { // 如果想省略的话 就用_
		fmt.Println(i, v)
	}

	//这个就是说把数组传递进去的时候 进行了拷贝 所以这里的数组是值类型
	//go语言中参数的 传递是值传递
	//go语言一般不直接使用数组 使用切片
	printArray(&arr3)

	fmt.Println("printArray3")
	fmt.Println(arr3)

}
