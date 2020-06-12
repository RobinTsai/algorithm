/* 判断一个二叉树树是否为平衡的
 * 平衡二叉树的定义：任意节点左右子树的高度差都小于等于 1
 * 方法（递归套路）：从左数上拿到左数的信息，从右树上拿到右树的信息
 * 然后整合生成自己的信息
 * 信息有：树的深度、是否为平衡二叉树
 */
package main

import "fmt"

func main() {
	head := getOneTree()
	result := isBalanceTree(head)
	fmt.Println(result)
}
func isBalanceTree(head *Node) bool {
	_, isBalance := balanceInfo(head)
	return isBalance
}

// balanceInfo 返回树的深度和是否为平衡树
func balanceInfo(cur *Node) (int, bool) {
	if cur == nil {
		return 0, true
	}
	leftDepth, isLeftBalance := balanceInfo(cur.Left)
	rightDepth, isRightBalance := balanceInfo(cur.Right)

	if absIntSub(leftDepth, rightDepth) <= 1 {
		return maxInt(leftDepth, rightDepth) + 1, isLeftBalance && isRightBalance
	}

	return maxInt(leftDepth, rightDepth) + 1, false
}

func absIntSub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
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
