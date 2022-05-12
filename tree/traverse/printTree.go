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

type HorizonTree Node

// Print 第二次盲写
func (n *HorizonTree) Print(depth int, flag direction) {
	if n == nil {
		return
	}
	(*HorizonTree)(n.Right).Print(depth+1, right)
	fmt.Print(strings.Repeat("  ", depth), flag, n.Value, "\n")
	(*HorizonTree)(n.Left).Print(depth+1, left)
}

// PrintAsTree 按树形打印输出（根节点在左）
func (n *Node) PrintAsTree() {
	n.printAsTreeNode(0, none)
}

func (n *Node) printAsTreeNode(depth int, flag direction) {
	if n == nil {
		return
	}
	n.Right.printAsTreeNode(depth+1, right)

	fmt.Printf("%s%s%d\n", strings.Repeat("   ", depth), flag, n.Value)
	n.Left.printAsTreeNode(depth+1, left)
}
