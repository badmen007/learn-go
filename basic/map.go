package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "xz",
		"age":     "12",
		"site":    "learn",
		"quality": "good",
	}
	fmt.Println(m)

	m2 := make(map[string]int) // m2 === empty map
	fmt.Println(m2)

	var m3 map[string]int // m3 === nil 可以参与运算 而不像 null
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for k, v := range m { // 每次出来的顺序还不一样
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"] // 取map中不存在的键的时候  拿到的是空串 怎么判断key在不在里面呢
	fmt.Println(courseName, ok)   // ok 布尔值 表示的是key 是不是存在

}
