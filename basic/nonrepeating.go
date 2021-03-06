package main

import "fmt"

func nonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}
	return maxLength
}

func main() {
	fmt.Println(
		nonRepeatingSubStr("abcabcbb"))
}
