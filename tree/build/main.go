package main

import (
	"fmt"
)

func main() {
	pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	mid := []int{4, 7, 2, 1, 5, 3, 8, 6}
	head := BuildFromPreAndMid(pre, mid)

	PreOrder(head)
	fmt.Println()
	MidOrder(head)
	fmt.Println()
	PostOrder(head)
}

// BuildFromPreAndMid 从先序和中序创建树
// 由性质来看，先序的第一个元素肯定是头结点，在中序中头结点将信息分成左树和右树。根据这点来做
func BuildFromPreAndMid(pre, mid []int) *Node {
	if len(mid) == 0 {
		return nil
	}
	headVal := pre[0]
	head := &Node{Value: headVal}

	leftVals := make([]int, 0)
	rightVals := make([]int, 0)
	leftMap := make(map[int]struct{})
	for i := 0; i < len(mid); i++ {
		if headVal == mid[i] {
			leftVals = mid[:i]
			rightVals = mid[i+1:]
			break
		}
		leftMap[mid[i]] = struct{}{}
	}

	leftPre := make([]int, 0)
	rightPre := make([]int, 0)
	for _, i2 := range pre[1:] {
		if _, ok := leftMap[i2]; ok {
			leftPre = append(leftPre, i2)
		} else {
			rightPre = append(rightPre, i2)
		}
	}
	head.Left = BuildFromPreAndMid(leftPre, leftVals)
	head.Right = BuildFromPreAndMid(rightPre, rightVals)

	return head
}

// Node 二叉树的结构
type Node struct {
	Value int
	Left  *Node
	Right *Node
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

// PreOrder 递归，先序，不用队列
func PreOrder(head *Node) {
	if head == nil {
		return
	}

	fmt.Print(head.Value, " ") // 先序
	PreOrder(head.Left)
	PreOrder(head.Right)
}

// MidOrder 递归，中序，不用队列
func MidOrder(head *Node) {
	if head == nil {
		return
	}

	MidOrder(head.Left)
	fmt.Print(head.Value, " ") // 中序
	MidOrder(head.Right)
}

// PostOrder 递归，后序，不用队列
func PostOrder(head *Node) {
	if head == nil {
		return
	}

	PostOrder(head.Left)
	PostOrder(head.Right)
	fmt.Print(head.Value, " ") // 后序
}
