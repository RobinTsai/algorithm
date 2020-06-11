package main

import "fmt"

/* 二叉树的递归套路：1. 输出二叉树某个节点的后继节点
 * 后继节点：按中序排序后排在某节点后面的节点都是此节点的后继节点
 *
 * 普通方式可以用中序遍历的方式，O(N) 的复杂度实现
 * 特殊的方式需要有特殊的结构（Parent 指向），可以实现 O(k) 的复杂度，k 为后继节点个数
 */

func main() {
	head, target := getSpecialTreeWithTarget()

	fmt.Println("O(n) 普通中序遍历：")
	SuccessorsFuncOne(head, target)

	fmt.Println("\n\nO(k) 的方式一：")
	SuccessorsFuncTwo(target)

	fmt.Println("\n\nO(k) 的方式二：")
	SuccessorsFuncThree(target)
}

// 特殊二叉树（包含 Parent 指向）的结构
type NodeWithParent struct {
	Value  int
	Left   *NodeWithParent
	Right  *NodeWithParent
	Parent *NodeWithParent
}

/* ----------------------- 方法三 O(k) ----------------------------- */
// 每次直接找到临近的后继节点，并输出，直到没有临近的后继
func SuccessorsFuncThree(cur *NodeWithParent) {
	for nextSuccessor(cur) != nil {
		cur = nextSuccessor(cur)
		fmt.Print(cur.Value, " ")
	}
}

// 返回 (cur 的) 临近的后继
// 方法是先找右树，如果有右树，则直接到右树的最左节点（此节点是临近后继）
// 如果没有右树，那么它的临近后继是在父节点中，且此节点所在子树是父节点的左分支（cur 是作为子树的最右）
func nextSuccessor(cur *NodeWithParent) *NodeWithParent {
	if cur == nil {
		return nil
	}
	if cur.Right != nil { // 有右子树
		return directToTheLeftest(cur.Right)
	}
	// 无右子树，那么向上攀升找本节点的临近后继
	for cur.Parent != nil && cur.Parent.Right == cur { // 若自己是父节点的右子节点，继续向上找，找到不是或没有父节点
		cur = cur.Parent // 注意，当找到时，是没有赋这个值的
	}
	return cur.Parent // 注意: 最后找到的临近后继不是 cur, 而还是 cur.Parent
}

// 一直找到最左节点
// 因为此子树的最左节点是上一个的后继节点
func directToTheLeftest(cur *NodeWithParent) *NodeWithParent {
	if cur.Left == nil {
		return cur
	}
	return directToTheLeftest(cur.Left)
}

/* ----------------------- 方法二 O(k)  ----------------------------- */
// SuccessorsFuncTwo 每一次调用输入 cur 节点
// 按 cur 输出所有右树节点（部分后继）
// 然后返回的是下一个既是后继，又是 parent 的节点
// 然后递归调用 SuccessorsFuncTwo... cur 为上一次的返回值 parent
func SuccessorsFuncTwo(cur *NodeWithParent) *NodeWithParent {
	if cur == nil { // 直到最后 cur 为 nil 结束
		return nil
	}

	midOrderFindSubTree(cur.Right) // 输出此节点右树的所有子节点
	nextCur := findNextParent(cur) // 从此节点向上找后继父节点
	if nextCur != nil {
		fmt.Print(nextCur.Value, " ")
	}
	return SuccessorsFuncTwo(nextCur) // 循环递归调用
}

// cur 祖先中，按中序遍历的下一个节点
func findNextParent(cur *NodeWithParent) *NodeWithParent {
	if cur == nil {
		return nil
	}
	nextCur := cur.Parent
	if nextCur != nil && nextCur.Right == cur {
		return findNextParent(nextCur)
	}
	return nextCur
}

// 一个完整的中序遍历，但只用于查找右树的子树
func midOrderFindSubTree(head *NodeWithParent) {
	if head == nil {
		return
	}
	midOrderFindSubTree(head.Left)
	fmt.Print(head.Value, " ")
	midOrderFindSubTree(head.Right)
}

/* ----------------------- 中序遍历的方法 O(n) ----------------------------- */
func SuccessorsFuncOne(head, target *NodeWithParent) {
	midOrderOutputSucceed(head, target)
	return
}

// 中序遍历当找到当前点后，后面的节点都为后继节点，此时设为 true
var isSuccessor bool

// 普通的递归中序遍历实现（不使用特殊字段 Parent ）
func midOrderOutputSucceed(head, target *NodeWithParent) {
	if head == nil {
		return
	}

	midOrderOutputSucceed(head.Left, target)
	if isSuccessor {
		fmt.Print(head.Value, " ") // 中序
	}
	if target == head { // 这里就是中序的位置，所以要将状态的变化加在这里
		isSuccessor = true
	}
	midOrderOutputSucceed(head.Right, target)
}

// --------------- 辅助函数 --------------
// 返回一个二叉树示例
//       A         1
//    B    C    2    3
//   D E  F G  4 5  6 7
func getSpecialTreeWithTarget() (*NodeWithParent, *NodeWithParent) {
	a := &NodeWithParent{Value: 1}
	b := &NodeWithParent{Value: 2}
	c := &NodeWithParent{Value: 3}
	d := &NodeWithParent{Value: 4}
	e := &NodeWithParent{Value: 5}
	f := &NodeWithParent{Value: 6}
	g := &NodeWithParent{Value: 7}

	a.Left = b
	b.Parent = a

	a.Right = c
	c.Parent = a

	b.Left = d
	d.Parent = b

	b.Right = e
	e.Parent = b

	c.Left = f
	f.Parent = c

	c.Right = g
	g.Parent = c

	return a, e
}
