/*

 */
package share

import (
	"fmt"
	"strings"
)

// 队列元素接口
type Element interface {
	String() string // 只需要一个打印输出的函数即可
}

// 队列提供的方法（这里只是为了方便查看，忽略规范性）
type Queue interface {
	Push(e Element) bool   // 入列
	Poll() (Element, bool) // 出列
	Size() int             // 返回队列大小
	Length() int           // 返回队列长度
	String() string        // 方便打印队列内容用于调试
}

// 一个简单的队列
type SimpleQueue struct {
	pushIdx int
	pollIdx int
	values  []Element
	size    int
}

// NewQueue 初始化一个队列
// size 个元素用 size+1 个空间来存储，多一个来标记空或满
// (用一个辅助变量 length 的话可以省去这一个空间)
func NewQueue(size int) *SimpleQueue {
	if size <= 0 {
		size = 1024 // default size
	}
	return &SimpleQueue{
		pushIdx: 0,                       // 指向应该弹出的点
		pollIdx: 0,                       // 指向下一个应该插入数据的点
		size:    size,                    //
		values:  make([]Element, size+1), // size 大小需要申请 size+1 的空间，多一个用来标记空或满
	}
}

func (queue *SimpleQueue) Push(e Element) bool {
	nextPushIdx := queue.getNextIdx(queue.pushIdx)
	if nextPushIdx == queue.pollIdx {
		// fmt.Println("queue is full")
		return false // 满
	}
	queue.values[queue.pushIdx] = e
	queue.pushIdx = nextPushIdx
	return true
}

func (queue *SimpleQueue) Poll() (Element, bool) {
	if queue.pushIdx == queue.pollIdx {
		// fmt.Println("queue is empty")
		return nil, false // 空
	}
	e := queue.values[queue.pollIdx]
	queue.pollIdx = queue.getNextIdx(queue.pollIdx)
	return e, true
}

func (queue *SimpleQueue) Size() int {
	return queue.size
}

func (queue *SimpleQueue) Length() int {
	if queue.pushIdx > queue.pollIdx {
		return queue.pushIdx - queue.pollIdx
	}
	return queue.pollIdx - queue.pushIdx
}

func (queue *SimpleQueue) getNextIdx(thisIdx int) int {
	next := thisIdx + 1
	if next > queue.size { // 或 next == queue.size+1
		next = 0
	}
	return next
}

func (queue SimpleQueue) String() string {
	valStr := ""
	for i := queue.pollIdx; i < queue.pushIdx; i++ {
		valStr = fmt.Sprintf("%s %v", valStr, queue.values[i])
	}
	return fmt.Sprintf("queue: [%s]", strings.Trim(valStr, " "))
}
