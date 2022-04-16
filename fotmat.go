package main

import "fmt"

func main() {
	name := "xz"
	age := 14
	// 这个API 是返回字符串的， 所以大多数我们都用这个 格式化 包含格式化输出
	str := fmt.Sprintf("hello, I am %s, I am %d years old", name, age)
	println(str)
}
