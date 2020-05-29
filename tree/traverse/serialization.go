package main

import "fmt"

/* 序列化反序列化转换
 * 序列化反序列化是为了方便持久化或者进行转移
 * 它们是一对相反的过程，所以规则必须对应
 * 注意对于操作的二叉树并不一定是完全二叉树
 * 所以为了标记树的叶子节点，要记录一下叶子的终止位置（叶子下为 nil 的节点）
 */
func Serialization() {
	head := getOneTree()
	fmt.Print("原树先序：")
	PreOrder(head) // 按先序输出树
	fmt.Println()

	// 序列化到 array 中
	arr := MarshalBinaryTreeToArray(head)

	fmt.Println("转成数组：", arr)

	// 反序列化
	newHead := UnmarshalArrayToBinaryTree(arr)
	fmt.Print("从数组恢复：")
	PreOrder(newHead) // 反序列化后输出树
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

// 注意这里的 nodes 需要用切片指针传入
// 参考：https://github.com/RobinTsai/Go-Questions/blob/master/数组和切片/切片作为函数参数.md
// 猜想究其原因可能是这样：
//   因为 slice 相当于一个结构体，包含 len, cap, array 三部分
//   我们常说在引用传递，应该就是指底层共用同一个数组这个特性
//   但它的 len 应该是值传递不会改变，所以必须用显示的指针符号
func PreOrderPushToArray(head *Node, nodes *[]*Node) {
	*nodes = append(*nodes, head)
	if head == nil {
		return
	}

	PreOrderPushToArray(head.Left, nodes)
	PreOrderPushToArray(head.Right, nodes)
}

// arr 是一个二叉树经过先序序列化后的数组
func UnmarshalArrayToBinaryTree(arr []int) *Node {
	cur := 0
	head := joinNode(nil, arr, &cur)
	return head
}

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
