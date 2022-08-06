package main

import (
	"fmt"
	"math"
)

// 常量的定义 常量不需要大写  因为在go语言中大写有特殊的含义
func consts() {
	//const filename = "abc.txt"
	//const a, b = 3, 4
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

//枚举类型
func enums() {
	const (
		cpp = iota // 表示自增 后面不用去赋值了
		java
		python
		golang
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	consts()

	enums()
}
