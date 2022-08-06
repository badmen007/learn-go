package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// 可以返回多个值
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	//拿到函数的名字
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 传递可变参数的列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

// 交换两个数  值传递和引用传递?
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	if result, error := eval(3, 4, "x"); error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))
	sum(1, 2, 3, 4, 5, 6)

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
