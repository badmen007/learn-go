package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:]", s1)
	s2 := arr[:]
	fmt.Println("arr[:]", s2)

	fmt.Println("after updateSlice s1")
	updateSlice(s1)
	fmt.Println(s1)

	fmt.Println("after updateSlice s2")
	updateSlice(s2)
	fmt.Println(s2)

	fmt.Println("Reslice")
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1=%v, len=%d, cap=%d \n", s1, len(s1), cap(s1)) // 明明没有了为什么还能取到？
	fmt.Printf("s2=%v, len=%d, cap=%d \n", s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5=", s3, s4, s5)
	// s4 and s5 no longer view arr
	// 添加元素的时候如果超越cap, 系统会重新分配更大的底层数组 原来的数组如果没人用的话 就会被垃圾回收掉
	// 由于是值传递的关系 必须接受append的值
	fmt.Println("arr=", arr)
}
