package main

import (
	"fmt"
	"strings"
)

type direction string

const (
	none  direction = ""
	left  direction = "∧" // 表示自己为左节点
	right direction = "∨" // 表示自己为右节点
)

// 按树形打印输出（根节点在左）
func (head *Node) PrintAsTree() {
	head.printAsTreeNode(0, none)
}

func (node *Node) printAsTreeNode(depth int, flag direction) {
	if node == nil {
		return
	}
	node.Right.printAsTreeNode(depth+1, right)

	fmt.Printf("%s%s%d\n", strings.Repeat("   ", depth), flag, node.Value)
	node.Left.printAsTreeNode(depth+1, left)
}
