package main

import "fmt"

type traverseByDepth struct{}
type printEveryBranch traverseByDepth

// PrintEveryBranch 输出二叉树的所有完整分支
var PrintEveryBranch printEveryBranch

// 递归：recursion
// 迭代：iteration
// 回溯：backtrace

// ByRecursionWithBacktrace 使用递归的方式进行回溯
func (*printEveryBranch) ByRecursionWithBacktrace(head *Node) {
	if head == nil {
		return
	}
	prefix := new([]*Node)
	*prefix = make([]*Node, 0, 10)

	var recursion func(prefix *[]*Node, head *Node) // 函数中用了递归，所以需要先声明
	recursion = func(prefix *[]*Node, head *Node) { // 递归函数
		*prefix = append(*prefix, head)
		if head == nil {
			return
		}
		if head.Left == nil && head.Right == nil {
			for _, node := range *prefix {
				fmt.Print(node.Value, " ")
			}
			fmt.Println()
			return
		}

		recursion(prefix, head.Left)
		*prefix = (*prefix)[:len(*prefix)-1] // Pop

		recursion(prefix, head.Right)
		*prefix = (*prefix)[:len(*prefix)-1] // Pop
	}

	recursion(prefix, head)
}

// ByIterationWithBacktrace 使用迭代的方式进行回溯
// --- (使用迭代的方式比较考验逻辑能力) ---
// 对于左节点的话，直接在 push 某个结点的时候，就把所有左侧结点 push 进去（相当于左侧边拉平，只判断右侧点）
// 对于右结点要有个标记来标记是从右结点回来还是要进右结点
// 然后对于处理结点要统一处理，这样可以避免一些特殊边界判断:这边所有的判断右结点都从 stack 中取
func (*printEveryBranch) ByIterationWithBacktrace(head *Node) {
	if head == nil {
		return
	}

	// 定义输出函数
	printx := func(stack []*Node) {
		for _, node := range stack {
			fmt.Print(node.Value, " ")
		}
		fmt.Println()
	}
	// 当 push 一个结点的时候，把所以的左侧结点都 push 进去
	pushWithLeft := func(stack []*Node, head *Node) []*Node {
		for head != nil {
			stack = append(stack, head)
			head = head.Left
		}
		return stack
	}

	stack := make([]*Node, 0, 10)
	stack = pushWithLeft(stack, head) // 约定: 在 push 某个点的时候把所有左侧点都 push 进去
	var last *Node                    // 对于 right 只用栈没法判断是从 right 回来的情况，这种情况不能再重复 push，所以用 last 来判断

	for len(stack) > 0 {
		head = peek(stack)
		if head.Right == nil { // 右结点 nil，可能是叶子结点也可能不是。分析完弹出
			if head.Left == nil { // 如果左结点也为 Nil 说明是叶子，输出
				printx(stack)
			}
			last, stack = pop(stack)
			continue
		}
		if last == head.Right { // 右结点已经分析过，弹出并继续上一个结点分析
			last, stack = pop(stack)
			continue
		}

		head = head.Right // 第一次访问此 right 结点，加入分析流程
		stack = pushWithLeft(stack, head)
	}
}
