package main

import (
	"fmt"
	. "robintsai/algorithm/tree/share"
)

type Edges map[*Node]*Node

// TraverseByLayer_ConvertToEdges 按层遍历，迭代，普通数组方式
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

// TraverseByLayer_Queue 按层遍历（队列）
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
