// 面试官为了考察逻辑能力，会专门让你用最原始的方法按序遍历二叉树树
// 遍历树按 先序、中序、后序 依次变难，见 README.md 中先序、中序、后序的解释：
// 先序是第一次访问结点时输出（入栈时是第一次访问此结点，所以在入栈时输出是先序）
// 中序是第二次访问结点时输出（出栈时是第二次访问此节点，所以出栈时输出时中序，也可以时是为从左节点回来）
// 后序是第三次访问结点时输出（用第二个栈，从第二个栈中出来变成第三次访问。另外用从右结点回来的判断时需要加个标志）
package main

import "fmt"

// IterationPreOrMid 先序和中序可以参照此模板
// 差别就在于处理输出逻辑的位置不同
// 后序的话因为需要额外的辅助，所以参考
func IterationPreOrMid(head *Node, order string) {
	if head == nil {
		return
	}

	pre := func(val *Node) {
		if order == "pre" {
			fmt.Print(val, " ")
		}
	}
	mid := func(val *Node) {
		if order == "mid" {
			fmt.Print(val, " ")
		}
	}

	stack := make([]*Node, 0, 10)
	start := true
	for len(stack) > 0 || start || head != nil {
		start = false
		for head != nil { // 1
			pre(head) //
			stack = append(stack, head)
			head = head.Left
		}
		head, stack = stack[len(stack)-1], stack[:len(stack)-1] // 2
		mid(head)
		head = head.Right // 3
	}
	fmt.Println()
}

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
		fmt.Print(head.Value, " ")              //
		head = head.Right                       // 右结点作为 head，继续下一轮（注意这里不判断是否 nil 都进入下一轮）
	}
	fmt.Println()
}

// RawPreTraverse 纯逻辑（非递归）实现先序遍历
// 关键点：也是让每个元素都经历入栈出栈，head.Left/Right 如果是 nil 也让他作为一次 head（只不过不会入栈），栈非空且head=nil 时从栈中取
// 入栈是第一次操作，所以在这里处理就是先序
// 出栈是第二次操作，所以在这里操作就是中序（head 是 nil 的时候就从栈中取，所以想法让左右结点都分析完后就设置 head nil）
// 第三次操作无法用栈来识别，所以需要引入另一个字段标记
// 先写内存循环，再写外层
func RawPreTraverse(head *Node) {
	if head == nil {
		return
	}
	stack := make([]*Node, 0, 10)
	start := true
	for len(stack) > 0 || // 主要条件
		start || head != nil { // 两个特殊情况条件（start 多余了，因为 head 首次不为 nil）

		start = false
		for head != nil {
			fmt.Print(head.Value, " ")
			stack = append(stack, head)
			head = head.Left
		}
		lastIdx := len(stack) - 1
		head, stack = stack[lastIdx], stack[:lastIdx]
		head = head.Right // 巧妙利用 nil 的 head，当 head == nil 时会进入下一次循环从 stack 中取 head
	}
	fmt.Println()
}

// RawPreTraverse2 出栈时候输出
// 这就要求入栈时先压右再压左，思路会更简单
func RawPreTraverse2(head *Node) {
	if head == nil {
		return
	}
	stack := make([]*Node, 0, 10)
	stack = append(stack, head)

	for len(stack) > 0 {
		head, stack = pop(stack)
		fmt.Print(head.Value, " ")

		if head.Right != nil {
			stack = append(stack, head.Right)
		}
		if head.Left != nil {
			stack = append(stack, head.Left)
		}
	}
	fmt.Println()
}

func pop(stack []*Node) (*Node, []*Node) {
	node := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return node, stack
}
func peek(stack []*Node) *Node {
	return stack[len(stack)-1]
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
		if last.Right == nil || last.Right == fromNode { // 从右结点回来或右结点为 nil
			fmt.Print(last.Value, " ")
			fromNode = last
			stack = stack[:len(stack)-1]
			continue
		}
		head = last.Right // 右结点重新在 for 中循环分析
	}
	fmt.Println()
}

// 因为栈只有入和出两个状态，可以在出栈时判断后序的访问，这就要求在分析完当前节点的右节点之后再出栈
// 总思路就是左节点入栈完后右节点也入栈，等右节点出栈分析完后（包含已出栈），输出当前节点
// 当问题在于怎么判断右节点分析完，判断方法是缓存一下上次弹出的节点 last
// 如果 last == peek.Right 则 peek 出栈并打印（peek 为栈顶节点元素）
// nil 不会入栈
func postOrder(root *Node) { // 再加一个分析
	if root == nil {
		return
	}
	stack := make([]*Node, 0)
	cur := root

	var last *Node
	for len(stack) > 0 || cur != nil {
		for cur != nil { // 所有左结点入栈，直到 nil
			stack = append(stack, cur)
			cur = cur.Left
		}
		peek := stack[len(stack)-1] // 先获取信息不出栈（这里不用考虑 nil，因为栈中没有 nil 值且一定能获取到）

		if peek.Right == nil || peek.Right == last { // 右节点分析完。条件：peek.right 是 nil 或着 last，则是第三次遍历到
			fmt.Print(peek.Val, " ")                                // 输出
			last, stack = stack[len(stack)-1], stack[:len(stack)-1] // 更新 last 和栈
			continue                                                // 如果下方有 else 此步可去掉
		} else { // 如果上方有 continue 此 else 可去掉
			// 如果有右节点，则更新当前值，进入 for 循环进行分析
			cur = peek.Right // 更新 cur
		}
	}
}

// RawPostTraverse_WithTwoStack 用两个栈实现后序输出
// 则需要先左后右入栈 1，然后先右后左出栈后入栈 2
// 最后出栈 2 的时候变成先左后右，即后序
func RawPostTraverse_WithTwoStack(head *Node) {
	if head == nil {
		return
	}
	stack := make([]*Node, 0, 10)  // 栈 1
	stack2 := make([]*Node, 0, 10) // 栈 2

	stack = append(stack, head)
	for len(stack) > 0 {
		head, stack = pop(stack)      // 先右后左出栈 1，紧接着入栈 2 即也是先右后左
		stack2 = append(stack2, head) // 先右后左入栈 2

		// 先左后右入栈 1
		if head.Left != nil {
			stack = append(stack, head.Left) // 不要变更 head，因为下方还用到它的 Right 的值
		}
		if head.Right != nil {
			stack = append(stack, head.Right)
		}
	}

	for len(stack2) > 0 { // 先左后右出栈 2
		head, stack2 = pop(stack2)
		fmt.Print(head, " ")
	}
	fmt.Println()
}
