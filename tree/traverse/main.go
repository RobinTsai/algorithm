/** 1. 二叉树的先序、中序、后序遍历（PreOrder, MidOrder, PostOrder）
 * 代码上体现就是：
 * 先序是递归之前输出 head（自己）
 * 中序是在遍历左节点后输出 head
 * 后序是在遍历左右节点后输出 head
 *
 * 2. 非递归/迭代方式实现先序、中序、后序（代码略）
 * 用栈来操作，先序步骤
 * a. head 压入栈
 * b. head 弹出栈，弹出即打印
 * c. 如果有右节点压入栈
 * d. 如果有左节点压入栈
 * ...
 * 后序与先序类型，需要先考虑左节点，再考虑右节点
 * 中序略难一些，是先将左侧所有的元素放入栈，弹出最后一个并打印，再分析栈顶元素的右侧
 */
package main

import (
	"fmt"
	. "robintsai/algorithm/tree/share"
)

// 二叉树的结构
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	head := getOneTree()

	// 迭代方式入列出列实现按层遍历
	TraverseByLayer_Queue(head)
	fmt.Println()

	fmt.Print("\n递归方式，先序，用队列，输出：")
	q1 := NewQueue(100)
	recursivePushToQueue_preOrder(head, q1)
	for cur_q1, ok := q1.Poll(); ok; cur_q1, ok = q1.Poll() {
		fmt.Print(" ", cur_q1)
	}

	fmt.Print("\n递归方式，中序，用队列，输出：")
	q2 := NewQueue(100)
	recursivePushToQueue_midOrder(head, q2)
	for cur_q2, ok := q2.Poll(); ok; cur_q2, ok = q2.Poll() {
		fmt.Print(" ", cur_q2)
	}

	fmt.Print("\n递归方式，后序，用队列，输出：")
	q3 := NewQueue(100)
	recursivePushToQueue_postOrder(head, q3)
	for cur_q3, ok := q3.Poll(); ok; cur_q3, ok = q3.Poll() {
		fmt.Print(" ", cur_q3)
	}

	fmt.Println()
	fmt.Print("\n递归先序，不用队列：")
	PreOrder(head)
	fmt.Println()

	fmt.Print("递归中序，不用队列：")
	MidOrder(head)
	fmt.Println()

	fmt.Print("递归后序，不用队列：")
	PostOrder(head)
	fmt.Println()

	fmt.Println("\n输出叶子节点和逆序的边集：")
	leaves, edges := TraverseByLayer_ConvertToEdges(head)
	fmt.Println("叶子节点：", leaves)
	fmt.Println("逆序边集：", edges)

}

// 递归，先序，不用队列
func PreOrder(head *Node) {
	if head == nil {
		return
	}

	fmt.Print(head.Value, " ") // 先序
	PreOrder(head.Left)
	PreOrder(head.Right)
}

// 递归，中序，不用队列
func MidOrder(head *Node) {
	if head == nil {
		return
	}

	MidOrder(head.Left)
	fmt.Print(head.Value, " ") // 中序
	MidOrder(head.Right)
}

// 递归，后序，不用队列
func PostOrder(head *Node) {
	if head == nil {
		return
	}

	PostOrder(head.Left)
	PostOrder(head.Right)
	fmt.Print(head.Value, " ") // 后序
}

type Edges map[*Node]*Node

// 按层遍历，迭代，普通数组方式
// 输入头节点
// 返回树的所有 叶子节点 和 反向边集
// （此函数的作用：从叶子节点开始按反向边集查，能查到每一个链路）
func TraverseByLayer_ConvertToEdges(head *Node) ([]*Node, Edges) {
	var edges Edges = Edges{head: nil} // 根节点的逆序边，可以有也可以没有
	leaves := []*Node{}

	layer := 0    // 层数
	maxWidth := 0 // 最大宽度（不完全二叉树的最大宽度可能是最后一层的宽度，也可能是倒数第二层的宽度）
	nextNodes := []*Node{head}
	for len(nextNodes) != 0 {
		layer++
		tempWidth := 0

		tempNextNodes := []*Node{}
		for _, node := range nextNodes {
			tempWidth++

			if node.Left != nil {
				edges[node.Left] = node
				tempNextNodes = append(tempNextNodes, node.Left)
			}
			if node.Right != nil {
				edges[node.Right] = node
				tempNextNodes = append(tempNextNodes, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				leaves = append(leaves, node)
			}
		}
		nextNodes = tempNextNodes
		if tempWidth > maxWidth {
			maxWidth = tempWidth
		}
	}

	fmt.Println("中间量，total layers:", layer)
	fmt.Println("中间量，max width:", maxWidth)

	return leaves, edges
}

func (n Node) String() string {
	return fmt.Sprintf("%d", n.Value)
}

// 按层遍历（队列）
// 迭代方式入列出列
func TraverseByLayer_Queue(head *Node) {
	q1 := NewQueue(100)
	q1.Push(head) // 要先将 head 节点设置进去（出列时打印并分析子节点）

	fmt.Print("迭代方式入列出列实现按层遍历：")
	for cur := iterateByQueue(q1); cur != nil; cur = iterateByQueue(q1) {
		fmt.Print(" ", cur)
	}
}

// 按层遍历（队列）
// 迭代方式，按层逐个入列再出列并输出返回
// 输入的队列中要有一个 head 节点
func iterateByQueue(q *SimpleQueue) *Node {
	cur, ok := q.Poll()
	if !ok {
		return nil
	}
	node := cur.(*Node)
	if node.Left != nil {
		q.Push(node.Left)
	}
	if node.Right != nil {
		q.Push(node.Right)
	}
	return node
}

// 递归方式，先序，用队列
func recursivePushToQueue_preOrder(head *Node, q *SimpleQueue) {
	q.Push(head)
	if head.Left != nil {
		recursivePushToQueue_preOrder(head.Left, q)
	}
	if head.Right != nil {
		recursivePushToQueue_preOrder(head.Right, q)
	}
}

// 递归方式，中序，用队列
func recursivePushToQueue_midOrder(head *Node, q *SimpleQueue) {
	// 中序方式加入队列
	if head.Left != nil {
		recursivePushToQueue_midOrder(head.Left, q)
	}
	q.Push(head)
	if head.Right != nil {
		recursivePushToQueue_midOrder(head.Right, q)
	}
}

// 递归方式，后序，用队列
func recursivePushToQueue_postOrder(head *Node, q *SimpleQueue) {
	// 后序加入队列
	if head.Left != nil {
		recursivePushToQueue_postOrder(head.Left, q)
	}
	if head.Right != nil {
		recursivePushToQueue_postOrder(head.Right, q)
	}
	q.Push(head)
}

// --------------- 辅助函数 --------------
// 返回一个二叉树示例
//       A         1
//    B    C    2    3
//   D E  F G  4 5  6 7
func getOneTree() *Node {
	a := &Node{Value: 1}
	b := &Node{Value: 2}
	c := &Node{Value: 3}
	d := &Node{Value: 4}
	e := &Node{Value: 5}
	f := &Node{Value: 6}
	g := &Node{Value: 7}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Left = f
	c.Right = g
	return a
}
