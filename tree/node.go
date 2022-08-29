package tree

import "fmt"

type Node struct { // 结构创建在堆上还是在栈上？ 不需要知道
	Value       int
	Left, Right *Node
}

// Print 就是给这个结构定义方法
func (node Node) Print() {
	fmt.Print(node.Value)
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

// Traverse 遍历
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse() // 这里不用判断这个left 是否存在
	node.Print()
	node.Right.Traverse()
}

func CreateTreeNode(value int) *Node {
	return &Node{Value: value} // 这个返回的是局部变量的地址
}
