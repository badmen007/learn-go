package main

import "fmt"

// 增加
func Add(s []int, index int, value int) []int {

	return s
}

// 删除
func Delete(s []int, index int) []int {
	return s
}

// 判断数组和切片的方法就是[]中有没有数字
func main() {
	s1 := []int{1, 2, 3, 4} // 初始化四个元素的切片
	fmt.Printf("s1: %v, len:%d, cap:%d", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4) // 创建了一个包含三个元，容量为4的切片
	fmt.Printf("s2: %v, len: %d, cap:%d", s2, len(s2), cap(s2))

	s2 = append(s2, 7) // 后面添加一个元素， 没有超出容量的限制的时候， 不会发生扩容
	// 0, 0, 0, 7
	fmt.Printf("s2: %v, len: %d, cap:%d", s2, len(s2), cap(s2))

	s2 = append(s2, 8)                                          //
	fmt.Printf("s2: %v, len: %d, cap:%d", s2, len(s2), cap(s2)) // 超出容量的时候会发生扩容的情况

	s3 := make([]int, 4) // 当只传一个参数，表示创建于给包含四个元素，容量也为4
	// 等价于 s3 := make([]int, 4, 4)
	fmt.Printf("s3: %v, len: %d, cap:%d", s3, len(s3), cap(s3))

	// 推荐写法 最佳实践 s1 := make([]type, 0, capacity)

	// 子切片和切片会不会相互影响 就抓住一点 他们是不是还共享数组？ 只需要看原本的切片 和 子切片 有没有扩容 扩容了就不共享了

}
