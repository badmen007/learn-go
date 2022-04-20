package main

import (
	"fmt"
	"io/ioutil"
)

// 1. if条件里可以赋值
// 2. if条件里赋值的变量作用域就在这个if语句里
func main() {
	const filename = "abc.txt"
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
