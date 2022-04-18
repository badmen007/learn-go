package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 在函数外面也能定义变量
// 包内变量
//var aa = 3
//var ss = "lkk"
//var bb = true

var (
	aa = 3
	ss = "kkk"
	bb = true
)

// 变量的定义
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, s, b)
}

// 变量的类型推断 编译器自己决定是什么类型
func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "def"
	var s = "abc"
	fmt.Println(a, b, c, d, s)
}

func variableShorter() {
	// 这种定义的方式只能在函数内部使用
	a, b, c, s := 3, 4, true, "def"
	b = 4
	fmt.Println(a, b, c, s)
}

// 欧拉公式
func euler() {
	//c := 3 + 4i
	//fmt.Println(complex.Abs(c))

	//fmt.Println(
	//	complex.Exp(1i*math.Pi) + 1)

	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) // 这里需要float64 传入的值和输出的值是float64 强制类型的转换
	fmt.Println(c)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
}
