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

func main() {
	head := buildExampleTree()
	longestDistance := getLongestDistance(head)
	fmt.Println(longestDistance)
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
//       A          1
//    B     C    2     3
//     E   F      5   6
//      D G        4 7
func buildExampleTree() *Node {
	a := &Node{Value: 1}
	b := &Node{Value: 2}
	c := &Node{Value: 3}
	d := &Node{Value: 4}
	e := &Node{Value: 5}
	f := &Node{Value: 6}
	g := &Node{Value: 7}

	a.Left = b
	a.Right = c
	b.Right = e
	e.Right = d
	c.Left = f
	f.Left = g
	return a
}
