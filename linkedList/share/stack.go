package share

import (
	"fmt"
	"math/rand"
)

type Element interface{}

type Stack struct {
	Elements []Element
	Size     int
	MaxLen   int
}

func NewEmptyStack(maxLen int) *Stack {
	s := &Stack{}
	s.Elements = make([]Element, maxLen)
	s.MaxLen = maxLen
	s.Size = 0
	return s
}

func (s *Stack) Pop() Element {
	if s.Size <= 0 {
		return nil
	}
	e := s.Elements[s.Size-1]
	s.Size--
	return e
}

func (s *Stack) Push(e Element) {
	if s.Size > s.MaxLen {
		panic("overflowed")
	}

	s.Elements[s.Size] = e
	s.Size++
}

// ----------------------
//         测试
// ----------------------

func (s Stack) String() string {
	str := "Top to bottom: "
	if s.Size == 0 {
		return "nil"
	}
	for s.Size > 1 {
		str += fmt.Sprintf("%v->", s.Pop())
	}
	str += fmt.Sprintf("%v", s.Pop())
	return str
}

// NewRandomIntStack 为了测试方法，stack 默认最大设为了 1024
// 大于此值可以创建成功，但不能再添加元素
func NewRandomIntStack(size int) *Stack {
	maxLen := 1024
	if size > maxLen {
		maxLen = size
	}
	s := NewEmptyStack(maxLen)

	for ; size > 0; size-- {
		val := rand.Int() % 10
		s.Push(val)
	}
	return s
}
