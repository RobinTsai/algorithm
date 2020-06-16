/* 返回二叉树中的 **最大二叉搜索子树** 的头节点
 * 搜索树：整个树上没有重复值，且左树的值<本节点<右树的值
 * 二叉树的递归套路：从左右子树上要信息，封装成自己节点的信息
 */
package main

import (
	"fmt"
)

// 1. 先确定本节点需要如何通过子树哪些信息进行计算
//   1.1 左右树是不是搜索树，若不是，本节点一定不是
//   1.2 左树的最大值 < 本节点值 < 右树的最小值
//   1.3 本节点树的大小 = 左节点树的大小 + 右节点树的大小 + 1

type bstInfo struct {
	isBST         bool // 当前节点的树是否为搜索树
	maxSubBSTSize int  // 当前节点下搜索树部分的最大长度
	min           int  // 搜索树部分的最小值
	max           int  // 搜索树部分的最大值
}

func main() {
	example := getOneTree()

	info := getBSTInfo(example)

	if info == nil {
		fmt.Println(0)
	} else {
		fmt.Println(info.maxSubBSTSize)
	}
}

func getBSTInfo(head *Node) *bstInfo {
	if head == nil {
		return nil
	}
	lInfo := getBSTInfo(head.Left)
	rInfo := getBSTInfo(head.Right)

	return buildCurInfo(head, lInfo, rInfo)
}

func buildCurInfo(cur *Node, lInfo, rInfo *bstInfo) *bstInfo {
	if lInfo == nil || rInfo == nil { // 有 nil 的值的
		return buildCurInfoIfContainsNil(cur, lInfo, rInfo)
	}
	if lInfo.isBST && rInfo.isBST { // 两个子树都是搜索树
		return buildCurInfoWithBothIsBst(cur, lInfo, rInfo)
	}
	if lInfo.isBST { // 左子树是搜索树
		lInfo.isBST = false
		return lInfo
	}
	if rInfo.isBST { // 右子树是搜索树
		rInfo.isBST = false
		return rInfo
	}
	// 两者都不是搜索树，取包含最大的
	bst := lInfo
	if rInfo.maxSubBSTSize > lInfo.maxSubBSTSize {
		bst = rInfo
	}
	return bst
}

func buildCurInfoIfContainsNil(cur *Node, lInfo, rInfo *bstInfo) *bstInfo {
	if lInfo == nil && rInfo == nil { // 无左右子节点
		return &bstInfo{isBST: true, maxSubBSTSize: 1, min: cur.Value, max: cur.Value}
	}
	if rInfo != nil { // 只有右节点
		if rInfo.isBST && cur.Value < rInfo.min {
			return &bstInfo{true, rInfo.maxSubBSTSize + 1, cur.Value, rInfo.max}
		}
		rInfo.isBST = false
		return rInfo
	}

	if lInfo.isBST && cur.Value > lInfo.max {
		return &bstInfo{true, lInfo.maxSubBSTSize + 1, lInfo.min, cur.Value}
	}
	lInfo.isBST = false
	return lInfo
}

func buildCurInfoWithBothIsBst(cur *Node, lInfo, rInfo *bstInfo) *bstInfo {
	if lInfo.max < cur.Value && cur.Value < rInfo.min {
		return &bstInfo{
			isBST:         true,
			maxSubBSTSize: lInfo.maxSubBSTSize + rInfo.maxSubBSTSize + 1,
			min:           lInfo.min,
			max:           rInfo.max,
		}
	}

	maxOne := lInfo
	if lInfo.maxSubBSTSize < rInfo.maxSubBSTSize {
		maxOne = rInfo
	}
	maxOne.isBST = false
	return maxOne
}

// --------------- 辅助函数 --------------
// 返回一个二叉树示例
//       A         4
//    B    C    2    5
//   D E  F G  1 3  6 7
func getOneTree() *Node {
	a := &Node{Value: 4}
	b := &Node{Value: 2}
	c := &Node{Value: 5}
	d := &Node{Value: 1}
	e := &Node{Value: 3}
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

type Node struct {
	Value int
	Left  *Node
	Right *Node
}
