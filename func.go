package main

import "fmt"

// 多个参数 多个返回值 参数有名字 但是返回值没有
func name(name string, age int) (string, int) {
	return name, age
}

func Func1(name string) string {
	return "hello, " + name
}

//多个参数具有相同类型的放在一起可以只写一次
func Func2(a, b, c string, bcd int, p string) (d, e int, g string) {
	d = 15
	e = 16
	g = "nihao"
	return
}

//不定参数的问题
func Func3(a string, b string, names ...string) {

}
func main() {
	fmt.Println(name("alalla", 20))
	fmt.Println(Func1("xuezhuang"))
	fmt.Println(Func2("lll", "aaa", "dddd", 12, "ccc"))
}
