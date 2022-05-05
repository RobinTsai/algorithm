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

// Node 二叉树的结构
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	head := getOneTree()

	IterationPreOrMid(head, "pre")
	IterationPreOrMid(head, "mid")
	return
	PrintEveryBranch.ByRecursionWithBacktrace(head)

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

	fmt.Println("\n 序列化和反序列化：")
	Serialization()

	fmt.Println("\n\n按树结构输出：")
	head.PrintAsTree()
}

func (n Node) String() string {
	return fmt.Sprintf("%d", n.Value)
}

// --------------- 辅助函数 --------------
// 返回一个二叉树示例
//       A         1
//     B    C     2    3
//   D  E  F G  4  5  6 7
//    H          8
func getOneTree() *Node {
	a := &Node{Value: 1}
	b := &Node{Value: 2}
	c := &Node{Value: 3}
	d := &Node{Value: 4}
	e := &Node{Value: 5}
	f := &Node{Value: 6}
	g := &Node{Value: 7}
	h := &Node{Value: 8}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Left = f
	c.Right = g
	d.Right = h
	return a
}
