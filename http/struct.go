package main

import "fmt"

type ToyDuck struct {
	Color string
	Price uint64
}

func (t *ToyDuck) Swim() {
	fmt.Printf("门前一条河， 游过一群鸭")
}

func main() {
	// duck1是 *ToyDuck
	duck1 := &ToyDuck{}
	duck1.Swim()

	duck2 := ToyDuck{}
	duck2.Swim()

	// duck3是 *ToyDuck
	duck3 := new(ToyDuck)
	duck3.Swim()

	// 当你声明这样的时候 go就帮你分配好
	// 不用担心空指针的问题，以为它压根就不是指针
	var duck4 ToyDuck
	duck4.Swim()

	var duck5 *ToyDuck
	duck5.Swim()

	// 赋值 初始化字段名字赋值
	duck6 := ToyDuck{
		Color: "蓝色",
		Price: 100,
	}
	duck6.Swim()
}
