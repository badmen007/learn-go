package main

import "fmt"

func ForLoop() {
	arr := []int{9, 8, 7, 6}
	index := 0
	for {
		if index == 3 {
			break // 跳出循环
		}
		fmt.Printf("#{index} => #{arr[index]}\n")
		index++
	}
	fmt.Println("for loop end \n ", arr)
}

func ForI() {
	arr := []int{9, 8, 7, 6}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d => %d \n", i, arr[i])
	}
	fmt.Println("for loop end \n ")
}

func ForR() {
	arr := []int{9, 8, 7, 6}

	for index, value := range arr {
		fmt.Printf("%d => %d\n", index, value)
	}

	for _, value := range arr {
		fmt.Printf("%d => %d\n", value)
	}
}

func main() {
	ForLoop()
	ForI()
	ForR()
}
