package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] =", arr[2:6])
	s1 := arr[2:]
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:6] =", arr[:6])
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1) // s1变了 arr也变了  slice是对底层数组的一个view
	fmt.Println(arr)

	// slice 可以向后扩展 但是不能向前扩展 向后扩展不能超过底层的数组 cap(s)

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)                  // 必须接受append的返回值  超越了会重新分配一个更大的数组
	fmt.Println("s3, s4, s5", s3, s4, s5) // s4 s5 不再是arr的 view
	fmt.Println("arr=", arr)

}
