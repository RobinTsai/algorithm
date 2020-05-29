package main

func main() {
	head := getOneTree()
	PreOrder(head)
	fmt.Println()

	arr := MarshalBinaryTreeToArray(head)
	fmt.Println(arr)

	newHead := UnmarshalArrayToBinaryTree(arr)

	PreOrder(newHead)  
}

func MarshalBinaryTreeToArray(head *Node) []int {
	nodes := make([]*Node, 0, 100)
	PreOrderPushToArray(head, &nodes)

	nodeValues := []int{} // 序列化为数组，-1 代表 nil
	for _, node := range nodes {
		if node == nil {
			nodeValues = append(nodeValues, -1)
		} else {
			nodeValues = append(nodeValues, node.Value)
		}
	}

	return nodeValues
}

// arr 是一个二叉树经过先序序列化后的数组
func UnmarshalArrayToBinaryTree(arr []int) *Node {
	cur := 0
	head := joinNode(nil, arr, &cur)
	return head
}

/*
joinNode() 错误的尝试：试图传入 head 指针，将 head 指针位置填充为正确的数据，而不再返回东西
发现,
规则一：当定义 head:=&Node{} 或 head:=new(Node) 后，两者都不等于 nil
	将 head 传入方法中，在方法内都是 head == nil
	此时 head 重新赋值会改变指针；取用 head.Value 会 panic
但当定义 head:=&Node{Value:0} 时，head 传入函数内不为 nil，可以去更新 head 中的值（但不能改变 head 指针指向）
*/
// 注意数组的下标需要是指针传递（或者用全局变量）
func joinNode(head *Node, arr []int, cur *int) *Node {
	head = generateNode(arr[*cur])
	// fmt.Println("cur:", *cur, "head:", head)
	if head == nil {
		return nil
	}
	*cur++
	head.Left = joinNode(head.Left, arr, cur)
	*cur++
	head.Right = joinNode(head.Right, arr, cur)
	return head
}

func generateNode(val int) *Node {
	if val == -1 {
		return nil
	}
	return &Node{Value: val}
}
