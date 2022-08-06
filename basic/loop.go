package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//十进制转换二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result //  strconv.Itoa 强制转换成string
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// 这里相当于while 省略起始条件 省略递增条件 go语言没有while  什么都不加的话那就是死循环
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1011 -> 1101
	)
	printFile("abc.txt")
}
