/* 获取二叉树的最长路径
 * 递归套路：
 * 1. 递归从左右子节点分别拿信息（各节点包含自己在内的最长距离，各节点的深度）
 * 2. 构造自己节点的信息（最长距离是 max(子节点最长距离, 左节点的深度+右节点深度+1)）
 * 3. 递归到结束
 */
package main

import (
	"fmt"
	"math"
)

// TheLongestDistance 作为 Receiver，获取树的最长路径
type TheLongestDistance struct {
	Head *Node
}

// Exec 方法：递归，从左右子节点上分别获取最大深度，并将自己加入后判断
func (r *TheLongestDistance) Exec() int {
	type depthInfo struct {
		curDepth    int // 最大深度
		maxDistance int // 最大长度
	}
	max := func(a int, bs ...int) int {
		for _, b := range bs {
			if b > a {
				a = b
			}
		}
		return a
	}

	var getDepthInfo func(head *Node) depthInfo
	getDepthInfo = func(head *Node) depthInfo {
		if head == nil {
			return depthInfo{0, 0}
		}

		li := getDepthInfo(head.Left)
		ri := getDepthInfo(head.Right)
		return depthInfo{
			curDepth:    max(li.curDepth, ri.curDepth) + 1,
			maxDistance: max(li.curDepth+ri.curDepth+1, li.maxDistance, ri.maxDistance),
		}
	}

	headInfo := getDepthInfo(r.Head)
	return headInfo.maxDistance
}

func main() {
	head := buildExampleTree()
	longestDistance := getLongestDistance(head)

	d2 := (&TheLongestDistance{head}).Exec()
	fmt.Println(longestDistance, d2)
}

// Node 二叉树的结点
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func getLongestDistance(head *Node) int {
	info := getLongestInfo(head)
	return info.LongestDistance
}

type longestInfo struct {
	Depth           int
	LongestDistance int
}

func getLongestInfo(head *Node) longestInfo {
	if head == nil {
		return longestInfo{0, 0}
	}
	infoL := getLongestInfo(head.Left)
	infoR := getLongestInfo(head.Right)

	longestDistance := getIntMax(infoL.Depth+infoR.Depth+1, infoL.LongestDistance, infoR.LongestDistance)
	curInfo := longestInfo{
		Depth:           getIntMax(infoL.Depth, infoR.Depth) + 1,
		LongestDistance: longestDistance,
	}
	fmt.Println(head.Value, curInfo)
	fmt.Println()
	return curInfo
}

func getIntMax(a ...int) int {
	max := math.MinInt64
	for _, val := range a {
		if val > max {
			max = val
		}
	}
	return max
}

// --------------- 辅助函数 --------------
// 返回一个二叉树示例
//     A
//    B
//   C E
//  F   D
// G
func buildExampleTree() *Node {
	a := &Node{Value: 1}
	b := &Node{Value: 2}
	c := &Node{Value: 3}
	d := &Node{Value: 4}
	e := &Node{Value: 5}
	f := &Node{Value: 6}
	g := &Node{Value: 7}

	a.Left = b
	b.Right = e
	b.Left = c
	e.Right = d
	c.Left = f
	f.Left = g
	return a
}
