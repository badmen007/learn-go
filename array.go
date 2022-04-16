package main

import "fmt"

// 1. 数组和其它的语言的数组差不多 语法是[cap]type
// 2. 初始化要指定长度(或者叫做容量)
// 3. 直接初始化
// 4. arr[i]的形式访问元素
// 5. len 和 cap 操作用于捕获数组的长度
func main() {
	// [cap]type
	a1 := [3]int{9, 8, 7}
	fmt.Printf("a1: %v, len: %d, cap: %d \n", a1, len(a1), cap(a1))

	var a2 [3]int
	fmt.Println(a2) //[0 0 0]
}
