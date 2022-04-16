package main

import "fmt"

func twoSum(nums []int, target int) (int, int) {
	// 暴力递归
	//for i := 0; i < len(nums); i++ {
	//	for j := 0; j < len(nums); j++ {
	//		if nums[i]+nums[j] == target {
	//			return i, j
	//		}
	//	}
	//}

	for i, v1 := range nums {
		for j, v2 := range nums {
			if v1+v2 == target {
				return i, j
			}
		}
	}

	return 0, 0
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
