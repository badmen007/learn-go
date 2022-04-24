package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 在go语言中函数是一等公民
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		_, r := div(a, b) // 下划线就表示 不想用
		return r, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s" + op)
	}
}

// 13 / 3 = 4 ...1
func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 函数作为一个参数传递
func apply(op func(int, int) int, a, b int) int {
	// 为了拿到函数的名字
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"%d, %d"+opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 这个就是你可以传递任意个参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4)) // 3的四次方

	fmt.Println(sum(1, 2, 4))
}
