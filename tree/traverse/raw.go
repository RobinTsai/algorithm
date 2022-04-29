// 面试官为了考察逻辑能力，会专门让你用最原始的方法按序遍历二叉树树
// 遍历树按 先序、中序、后序 依次变难，见 README.md 中先序、中序、后序的解释：
// 先序是第一次访问结点时输出
// 中序是第二次访问结点时输出（也即从左节点回来）
// 后序是第三次访问结点时输出（也即从右结点回来）
package main

import "fmt"

// RawMidTraverse 纯逻辑（非递归）实现中序遍历
// 关键点：让每个结点都有入栈出栈操作，这样相当于只需要考虑一个最基本的二叉树，不然会很难处理特殊情况。
// 这个时候当 head 为 nil 时从 stack 中取结点作为 head
func RawMidTraverse(head *Node) {
	stack := make([]*Node, 0, 10)
	if head == nil {
		return
	}
	start := true
	for len(stack) > 0 || // 核心条件
		start || // 特殊条件，首次时 len(stack) == 0
		head != nil { // 特殊条件，根结点出栈后 len(stack) == 0
		// 中序遍历中根结点会先于右结点输出，所以根节点出来时 head !=nil but len(stack) == 0
		start = false
		for head != nil { // 一直 push 左结点
			stack = append(stack, head)
			head = head.Left
		} // 直到 head == nil
		last := len(stack) - 1
		if last < 0 {
			return
		}
		head, stack = stack[last], stack[:last] // 从 stack 中拿出一个结点作为 head
		fmt.Print(head.Value, " ")              // 输出
		head = head.Right                       // 右结点作为 head，继续下一轮
	}
}

// RawPreTraverse 纯逻辑（非递归）实现先序遍历
// 关键点：也是让每个元素都经历入栈出栈，在入栈之前打印
// 先写内存循环，再写外层
func RawPreTraverse(head *Node) {
	if head == nil {
		return
	}
	stack := make([]*Node, 0, 10)
	start := true
	for len(stack) > 0 || // 主要条件
		start || head != nil { // 两个特殊情况条件
		start = false
		for head != nil {
			fmt.Println(head.Value)
			stack = append(stack, head)
			head = head.Left
		}
		lastIdx := len(stack) - 1
		head, stack = stack[lastIdx], stack[:lastIdx]
		head = head.Right
	}
}

// RawPostTraverse 就是 head 的左结点和右结点都访问完则输出本结点值
// 只用栈无法识别第三次访问结点，所以需要增加一个字段表示
// 这里增加字段 backFrom：如果遍历到当前结点是从右结点过来的，就可以证明此为第三次访问
func RawPostTraverse(head *Node) {
	if head == nil {
		return
	}
	stack := make([]*Node, 0, 10)
	var fromNode *Node // 用于标记从哪个结点回来，当从右结点后来的时候就完成了 head 左右子树的遍历
	start := true
	for len(stack) > 0 || start { // 2. 右结点
		start = false
		for head != nil { // 1. 左结点入栈
			stack = append(stack, head)
			head = head.Left
		}

		last := stack[len(stack)-1]
		if last.Right == nil ||
			fromNode == last.Right { // 从右结点回来或右结点为 nil
			fmt.Println(last.Value)
			fromNode = last
			stack = stack[:len(stack)-1]
			continue
		}
		head = last.Right // 右结点重新在 for 中循环分析
	}
}
