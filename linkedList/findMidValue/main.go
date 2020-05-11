package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cast"
)

// 找到链表奇数长度的中点、偶数长度的上中点
func (head *Node) getMidOrFirstMid() *Node {
	// 当本身、第二个元素、第三个元素为 nil 时
	// 即有 0个、1个、2个节点时，中点为 head
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	// 快慢指针
	tailPtr := head.Next.Next
	midPtr := head.Next
	// 快指针一次进 2 步，慢指针一次进 1 步
	for tailPtr.Next != nil && tailPtr.Next.Next != nil {
		midPtr = midPtr.Next
		tailPtr = tailPtr.Next.Next
	}

	return midPtr
}

// 找到链表奇数长度的中点、偶数长度的下中点
func (head *Node) getMidOrSecondMid() *Node {
	// 当本身、第二个元素为 nil 时
	// 即有 0个、1个节点时，中点为 head
	if head == nil || head.Next == nil {
		return head
	}
	// 快慢指针
	tailPtr := head.Next.Next
	midPtr := head.Next
	// 快指针一次进 2 步，慢指针一次进 1 步
	for tailPtr != nil && tailPtr.Next != nil {
		midPtr = midPtr.Next
		tailPtr = tailPtr.Next.Next
	}

	// 与 getMidOrFirstMid 只是多了此部分
	if tailPtr != nil {
		return midPtr.Next
	}

	return midPtr
}

// 找到链表奇数长度的中点前一个、偶数长度的上中点的前一个
func (head *Node) getPrevMidOrPrevFirstMid() *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	prevMidPtr := head
	tailPtr := head.Next.Next
	for tailPtr.Next != nil && tailPtr.Next.Next != nil {
		tailPtr = tailPtr.Next.Next
		prevMidPtr = prevMidPtr.Next
	}
	return prevMidPtr
}

// 找到链表奇数长度的中点前一个、偶数长度的下中点的后一个
func (head *Node) getPrevMidOrNextSecondMid() *Node {
	// 0 个、1 个、2 个时
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	// >= 3 个节点
	prevMidPtr := head
	tailPtr := head.Next.Next // 当前指在第三个位置

	for tailPtr.Next != nil && tailPtr.Next.Next != nil { // >= 3 个节点后
		tailPtr = tailPtr.Next.Next
		prevMidPtr = prevMidPtr.Next
	}

	if tailPtr.Next != nil { // 偶数时
		return prevMidPtr.Next.Next.Next
	}
	return prevMidPtr
}

// -------------------
// 以下为 debug 所需函数
// -------------------

// 主函数
func main() {
	times := 10
	for ; times > 0; times-- {
		if !test() {
			fmt.Println("Failed")
			break
		}
	}
	fmt.Println("Success")
}

func init() {
	rand.Seed(time.Now().Unix())
}

var debug = false

type Node struct {
	Value int
	Next  *Node
}

// 辅助函数，输出 linkedList 值
func (head *Node) ListString() (res string) {
	if head == nil {
		return "nil"
	}
	for head.Next != nil {
		res += cast.ToString(head.Value) + "->"
		head = head.Next
	}
	res += cast.ToString(head.Value)

	return res
}

// 用另一种方式去判断四个值是否取的正确
func (head *Node) assert(length int, firMid, secMid, prevMid, prevMidOrNextMid *Node) bool {
	switch length {
	case 0:
		return firMid == nil && secMid == nil && prevMid == nil && prevMidOrNextMid == nil
	case 1:
		return firMid == head && secMid == head &&
			prevMid == nil && prevMidOrNextMid == nil
	case 2:
		return firMid == head && secMid == head.Next &&
			prevMid == nil && prevMidOrNextMid == nil
	}

	firMidIdx := (length + 1) >> 1                        // (n+1)/2
	secMidIdx := ((length + 1) >> 1) + ((length + 1) & 1) // 奇 (n+1)/2，偶 (n+1)/2+1
	prevMidIdx := (length - 1) >> 1                       // (n-1)/2
	nextMidIdxOfEven := prevMidIdx                        // n 为奇数时为 (n-1)/2，为偶数时为 (n-1)/2+3
	if length&1 == 0 {
		nextMidIdxOfEven += 3
	}

	cur := head
	for i := 1; i <= length; i++ {
		switch i {
		case firMidIdx:
			if cur != firMid {
				return false
			}
		case secMidIdx:
			if cur != secMid {
				return false
			}
		case prevMidIdx:
			if cur != prevMid {
				return false
			}
		case nextMidIdxOfEven:
			if cur != prevMidOrNextMid {
				return false
			}
		}
		cur = cur.Next
	}

	return true
}

// 辅助函数，返回节点值，用于 debug
func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return cast.ToString(n.Value)
}

// 辅助函数，随机生成节点
func initRandomNode() *Node {
	return &Node{
		Value: rand.Int() % 10,
		Next:  nil,
	}
}

// 辅助函数，生成链表
func initLinkedList(count int) *Node {
	if count == 0 {
		return nil
	}

	head := initRandomNode()
	cur := head
	for count--; count > 0; count-- {
		cur.Next = initRandomNode()
		cur = cur.Next
	}
	return head
}

// 测试函数，测试结果是否正确
func test() bool {
	length := rand.Int() % 10
	head := initLinkedList(length)
	debugPrint("linked list is:", head.ListString())

	firMid := head.getMidOrFirstMid()
	if debug {
		debugPrint(firMid.String(), "is mid or 1st mid")
	}

	secMid := head.getMidOrSecondMid()
	debugPrint(secMid.String(), "is mid or 2nd mid")

	prevMid := head.getPrevMidOrPrevFirstMid()
	debugPrint(prevMid.String(), "is prev mid or prev 1st mid")

	prevMidOrNextMid := head.getPrevMidOrNextSecondMid()
	debugPrint(prevMidOrNextMid.String(), "is prev mid or next 2nd mid")

	return head.assert(length, firMid, secMid, prevMid, prevMidOrNextMid)
}

func debugPrint(args ...interface{}) {
	if !debug {
		return
	}
	fmt.Println(args...)
}
