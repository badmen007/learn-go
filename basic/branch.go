package main

import (
	"fmt"
	"io/ioutil"
)

/**
switch后面可以没有表达式
*/
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score)) // 中断程序的执行 报错
	case score < 60:
		g = "f"
	case score < 80:
		g = "c"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt" //  默认是根路径么  ？？
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	//fmt.Println(contents)  这个contents 除了block 这个变量就拿不到了

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(100),
	)
}
