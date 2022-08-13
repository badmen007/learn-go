package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我真的好喜欢你啊！" // utf8编码
	for _, b := range []byte(s) {
		fmt.Printf("%x ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d, %X)", i, ch)

	}
	fmt.Println()

	fmt.Println("Rune count", utf8.RuneCountInString(s)) // 解码出来数的

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("c ", ch)
	}

	for i, ch := range []rune(s) { // 这个是最上层的  rune相当于go 的char
		fmt.Printf("(%d %c)", i, ch)
	}

	// 使用len获得字节的长度
	// 使用[]byte获得字节

	// 字符串的其他操作 都放在strings中了
}
