package main

import (
	"fmt"
	. "robintsai/algorithm/tree/share"
)

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
