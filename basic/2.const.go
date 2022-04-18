package main

import (
	"fmt"
	"math"
)

func consts() {
	//const filename string = "abc.txt"
	//const a, b = 3, 4 // 不给类型的时候 就像是个纯文本
	// go语言的常量不用大写
	// 常量的定义
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 枚举值定义
func enums() {
	//const (
	//	cpp    = 0
	//	java   = 1
	//	python = 2
	//	golang = 3
	//)
	const (
		cpp = iota // 表示自增
		java
		python
		golang
		javascript
	)

	// b, kb, mb, gb, tb, pb 自增值
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	consts()
	enums()
}
