package main

import (
	"fmt"
	"imooc/tree"
)

// go mod init imooc 定义了下 不然的话不能自动引入自定义的包  看了好久  这个idea 的提示还是挺好的
func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)          // 这里的new是什么意思
	root.Right.Right = tree.CreateTreeNode(2) // 指针也能用点

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	root.Print() // 3

	root.Traverse() // 树的遍历

	// 值接收者 和 指针接收者
	//要改变内容必须使用指针接收者
	//结构过大也可以考虑使用指针接收者
	//一致性： 如有指针接收者，最好都用指针接收者

	// go没有构造函数
}
