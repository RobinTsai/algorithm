package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := &TreeNode{Val: 2}
	b := &TreeNode{Val: 3}
	c := &TreeNode{Val: 1}
	d := &TreeNode{Val: 3}
	e := &TreeNode{Val: 1}
	f := &TreeNode{Val: 1}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	count := pseudoPalindromicPaths(a)
	fmt.Println(count)
}

func isPresudoCycle(edges Edges, node *TreeNode) bool {
	cur := node
	valueCount := make(map[int]int, 0)
	for cur != nil {
		valueCount[cur.Val]++
		cur = edges[cur]
	}

	oddCount := 0
	for _, count := range valueCount {
		if count%2 == 1 {
			oddCount++
			if oddCount > 1 {
				return false
			}
		}
	}

	return true
}

type Edges map[*TreeNode]*TreeNode

func ConvertToEdges(root *TreeNode) ([]*TreeNode, Edges) {
	var edges Edges = Edges{root: nil}
	leaves := []*TreeNode{}

	nextNodes := []*TreeNode{root}
	for len(nextNodes) != 0 {
		tempNextNodes := []*TreeNode{}
		for _, node := range nextNodes {
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
	}
	return leaves, edges
}

func pseudoPalindromicPaths(root *TreeNode) int {
	leaves, edges := ConvertToEdges(root)

	count := 0
	for _, left := range leaves {
		if isPresudoCycle(edges, left) {
			count++
		}
	}
	return count
}
