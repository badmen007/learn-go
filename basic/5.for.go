package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 十进制转换成二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file) // 为什么要这一句

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func converToBin(m int) {
}

func main() {
	fmt.Println(
		convertToBin(5), //  101
		convertToBin(13),
	)
	printFile("abc.txt")
}
