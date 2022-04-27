package main

import "fmt"

// cap的扩充是不是一直是2倍的关系啊
func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d \n", s, len(s), cap(s))
}

func main() {
	var s []int // zero value  for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	// slice 的创建方式
	fmt.Println("Create slice")
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)

	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1) // 把s1拷贝到s2上
	printSlice(s2)

	fmt.Println("Deleting elements from slcie")
	s2 = append(s2[:3], s2[4:]...) // 删掉8
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	fmt.Println(s2)

	fmt.Println("Popping from backend")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	fmt.Println(s2)
}
