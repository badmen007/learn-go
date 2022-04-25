package main

import "fmt"

func swap(a, b int) (int, int) { // 加上返回值就换回来了
	return b, a
}

//func swap(a, b *int) { // 变成指针类型 但是这种方法应该是不推荐的 写法不是很好
//	*b, *a = *a, *b
//}
func main() {
	a, b := 3, 4
	//swap(&a, &b) 取地址 传递进去 看上去怪怪的
	a, b = swap(a, b)
	fmt.Println(a, b)
}
