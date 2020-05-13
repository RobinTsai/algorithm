package share

import (
	"math/rand"

	"github.com/spf13/cast"
)

type ISigleLinkedList interface {
	Insert(n *Node)
	Delete() *Node
}

type SigleLinkedList struct {
	Head *Node
	Tail *Node
}
type Node struct {
	Value interface{}
	Next  *Node
}

func (l *SigleLinkedList) Insert(value interface{}) {
	curNode := &Node{
		Value: value,
		Next:  nil,
	}
	if l.Head == nil {
		l.Head = curNode
		l.Tail = curNode
		return
	}
	l.Tail.Next = curNode
	l.Tail = l.Tail.Next
}

func (l *SigleLinkedList) Delete() *Node {
	del := l.Head
	l.Head = del.Next
	return del
}

// join 函数的 Node 可能是一个链表
func (l *SigleLinkedList) Join(n *Node) {
	l.Tail.Next = n
	cur := n
	for cur.Next != nil {
		cur = cur.Next
	}
	l.Tail = cur
}

// --------------------------------------------------------
//                  测试用函数
// --------------------------------------------------------

// 辅助函数，输出 linkedList 值
// 注意：若定义方法接收者为 指针 时，则在值上使用 fmt.Print() 不会调用此方法
// 但若定义为 值 接收者，则在指针上使用 fmt.Print() 仍可以调用此方法
func (l SigleLinkedList) String() (res string) {
	head := l.Head
	if head == nil {
		return "nil"
	}
	if head == head.Next {
		panic("Error: head == head.Next") // 在 String 函数中 panic 会被输出，而不是中止程序
	}
	for head.Next != nil {
		res += head.String() + "->"
		head = head.Next
	}
	res += head.String()

	return res
}

// 辅助函数，返回节点值，用于 debug
func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return cast.ToString(n.Value)
}

// 辅助函数，随机生成节点
func InitRandomIntNode() *Node {
	return &Node{
		Value: rand.Int() % 10,
		Next:  nil,
	}
}

// 生成链表
func InitRandomIntLinkedList(count int) *SigleLinkedList {
	if count == 0 {
		return &SigleLinkedList{}
	}

	head := InitRandomIntNode()
	cur := head
	for count--; count > 0; count-- {
		cur.Next = InitRandomIntNode()
		cur = cur.Next
	}
	return &SigleLinkedList{
		Head: head,
		Tail: cur,
	}
}
