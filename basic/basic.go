package main

import "fmt"

var aa = 3
var ss = "lll"

// 在函数的外面不能使用 := 声明变量

//也可以简写
var (
	cc = 2
	dd = "xz"
)

// 定义变量不赋初始值
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// 定义变量赋初始值
func variableInitialValue() {
	var c, d int // 可以赋值 也可以不赋值
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, s, b, c, d)
}

// 自己去推断类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 省略关键字var 变得更加简洁
func variableShorter() {
	a, b, c, s := 5, 6, true, "css"
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
