package main

import (
	"fmt"
	"math/rand"
	"robintsai/algorithm/linkedList/share"
)

func main() {
	share.Debug = true

	l := initPalindromeList(9)
	share.DebugPrint("initPalindromeList:", &l)

	if isPalindrome_1(l.Head) && // 方法一
		isPalindrome_2(l.Head) && // 方法二
		isPalindrome_3(l.Head) { // 方法三
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

// 将链表元素逐个加入栈
// 将栈元素与链表元素同时逐个取出，比较值
// 每个值进都一样，则为回文结构
func isPalindrome_1(head *share.Node) bool {
	if head == nil || head.Next == nil {
		return true
	}

	stack := share.NewEmptyStack(1024)
	cur := head
	for cur != nil {
		stack.Push(cur.Value)
		cur = cur.Next
	}

	listCur := head
	for listCur != nil {
		share.DebugPrint("Output per time:", listCur, stack.Pop())
		listCur = listCur.Next
	}
	return stack.Size == 0
}

// 快慢指针，慢指针总是快指针的（前）中点
// 慢指针在链表前半部分压栈
// 慢指针在链表后半部分时，与出栈元素逐个比较
// 到结束每个元素都相等则为回文
func isPalindrome_2(head *share.Node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	if head.Next.Next == nil {
		return head.Value == head.Next.Value
	}

	stack := share.NewEmptyStack(1024)
	stack.Push(head.Value)

	fastCur := head.Next.Next // 如果链表有可能有环，初始值要跳过 head
	slowCur := head.Next
	for fastCur.Next != nil && fastCur.Next.Next != nil {
		stack.Push(slowCur.Value)
		slowCur = slowCur.Next
		fastCur = fastCur.Next.Next
	} // 结束时，slowCur 指向（前）中点，stack 存了链表前一半无中点的数据

	share.DebugPrint(stack, ", slowCur:", slowCur)
	for slowCur.Next != nil {
		curValue := slowCur.Next.Value
		stackValue := stack.Pop()
		share.DebugPrint("Output per time:", curValue, stackValue)
		if curValue != stackValue {
			return false
		}
		slowCur = slowCur.Next
	}
	return stack.Size == 0 // 用了这样的方式是为了保险，以免程序出错
}

// 快慢指针，慢指针总是快指针的（前）中点
// 将链表后半部分反转
// 逐个比较两个前半部分链表和后链表
// 再将原链表还原
func isPalindrome_3(head *share.Node) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fastCur := head // 初始取 head 就可以了
	slowCur := head // 之前写的可能麻烦了一些
	for fastCur.Next != nil && fastCur.Next.Next != nil {
		fastCur = fastCur.Next.Next
		slowCur = slowCur.Next
	} // 结束时，slowCur 指向（前）中点，stack 存了链表前一半无中点的数据
	// ------------------------------------------------------------
	newHead := slowCur.Next
	slowCur.Next = nil
	firstHalfTail := slowCur
	var tempHead *share.Node

	// 后半部分反转，得到头部 newHead
	for newHead.Next != nil {
		slowCur = newHead.Next
		newHead.Next = tempHead
		tempHead = newHead
		newHead = slowCur
	}
	newHead.Next = tempHead

	share.DebugPrint("lastHalfPartConverted:", &share.SigleLinkedList{Head: newHead, Tail: nil}) // 后半部分反转后的链表
	// 注意，当原链表是奇数长度时，后反转的链比前半部分链少一个，但下面用最短的链去比较则没问题
	cur2 := newHead
	cur1 := head
	result := true
	for cur2 != nil {
		share.DebugPrint(cur1, cur2)
		if cur1.Value != cur2.Value {
			result = false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	share.DebugPrint("result:", result)

	// 恢复原链表后半部分
	cur3 := newHead
	tempHead = nil
	for cur3 != nil {
		cur3 = newHead.Next
		newHead.Next = tempHead
		tempHead = newHead
		newHead = cur3
	}
	newHead = tempHead

	share.DebugPrint("recoveredLastHalf:", &share.SigleLinkedList{Head: newHead, Tail: nil})
	firstHalfTail.Next = newHead // 拼接
	share.DebugPrint("recoveredRawList:", &share.SigleLinkedList{Head: head, Tail: nil})

	return result
}

// ------------------
// 辅助函数
// ------------------

func initPalindromeList(length int) share.SigleLinkedList {
	l := share.InitRandomIntLinkedList(0)
	arr := make([]int, length>>1) // 注意当用 make 时，第二个参数指定了 length，也就是已有两个值
	for i := 0; i < length>>1; i++ {
		val := rand.Int() % 10
		arr[i] = val
		l.Insert(val)
	}

	if length&1 > 0 {
		val := rand.Int() % 10
		l.Insert(val)
	}
	for j := len(arr); j > 0; j-- {
		l.Insert(arr[j-1])
	}

	return *l
}
