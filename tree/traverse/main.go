/** 1. 二叉树的先序、中序、后序遍历（PreOrder, MidOrder, PostOrder）
* 比如下面一个二叉树
*   A
* B   C
* 其中，B、C 代表 A 的左右子树（子树：包含它们的后代）
* 那么，
* 先序即 A->B->C
* 中序即 B->A->C
* 后序即 B->C->A
* 当用递归时，先序、中序、后序就很好理解：
* 我们规定遍历就是先分析一下当前节点，再分析它的左节点，再分析它的右节点
* 那么，上面的例子的遍历就是 A->B->B->B->A->C->C->C->A
* 重复三次B是因为后两次判断了B的左右子树后返回到B节点
* 所以先序、中序、后序分别就是某个节点第一次、第二次、第三次出现的序列
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

import "fmt"

// 二叉树的结构
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	head := getOneTree()

	fmt.Print("先序遍历：")
	PreOrder(head)
	fmt.Println()

	fmt.Print("中序遍历：")
	MidOrder(head)
	fmt.Println()

	fmt.Print("后序遍历：")
	PostOrder(head)
	fmt.Println()
}

// 先序遍历
func PreOrder(head *Node) {
	if head == nil {
		return
	}

	fmt.Print(head.Value, " ") // 先序
	PreOrder(head.Left)
	PreOrder(head.Right)
}

// 中序遍历
func MidOrder(head *Node) {
	if head == nil {
		return
	}

	MidOrder(head.Left)
	fmt.Print(head.Value, " ") // 中序
	MidOrder(head.Right)
}

// 后序遍历
func PostOrder(head *Node) {
	if head == nil {
		return
	}

	PostOrder(head.Left)
	PostOrder(head.Right)
	fmt.Print(head.Value, " ") // 后序
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
