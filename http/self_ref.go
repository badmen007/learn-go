package main

// 自己引用自己 里面使用指针就行了 其它的就不用管了
type Node struct {
	// 自引用只能使用指针
	//left  Node  无法估量大小 报错的原因就是无法计算大小
	//right Node

	// 为什么用指针就可以 因为指针的大小是固定的
	left  *Node
	right *Node
}

func main() {

}
