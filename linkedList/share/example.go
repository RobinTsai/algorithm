package share

import (
	"math/rand"
)

func LinkedListExample() {
	Debug = true

	list := InitRandomIntLinkedList(0)   // 初始化 0 个节点的链表
	for count := 5; count > 0; count-- { // 插入节点值
		list.Insert(count)
	}
	DebugPrint(list.String()) // 输出链表

	list2 := InitRandomIntLinkedList(5) // 用函数生成链表
	DebugPrint(list2.String())          // 输出链表

	list.Join(list2.Head)     // 拼接两个链表
	DebugPrint(list.String()) // 输出拼接的链表

	for list.Head != nil {
		del := list.Delete()     // 逐个取出
		DebugPrint(del.String()) // 输出
	}
}

func StackExample() {
	Debug = true

	size := rand.Int() % 10
	DebugPrint("init stack with size", size)

	s := NewRandomIntStack(size)
	DebugPrint(s.String())

	for s.Size > 0 {
		DebugPrint(s.Size, s.Pop())
	}
}
