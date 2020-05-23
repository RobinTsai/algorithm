/**
 * 将链表分段
 * 按一个给定值，将小于它的值放左边，大于它的值放右边，等于它的值放中间
 * 方法一用数组先分段，再重组链表（shardByValue_1）
 *     其中，将数组分段又提供了两种方法（splitArrByPivot_1、splitArrByPivot_1）
 *     这里面一定要注意临界值和特殊情况。现在就知道对数器的使用了。
 * 方法二直接在原链表上操作（TODO）
 */
package main

import (
	"fmt"
	"robintsai/algorithm/linkedList/share"
)

func main() {
	share.Debug = true

	l := share.InitRandomIntLinkedList(7)
	share.DebugPrint("raw list:", l)

	newList := shardByValue_1(l.Head, 3)
	fmt.Println(newList)
}

// 将节点放入数组
// 对数组进行排序
// 将排序的数组重组成链表
// 解释：这个返回掩盖了题目要求的分段的段的问题，但方法是按照题目要求过程做的，重点在方法上，而不是返回结果上
func shardByValue_1(head *share.Node, pivot int) *share.SigleLinkedList {
	arr := make([]*share.Node, 0, 100)
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	share.DebugPrint("push in arr:", arr)

	gteStart, gtStart := splitArrByPivot_1(arr, pivot)
	share.DebugPrint("sorted arr:", arr, gteStart, gtStart)

	var newHead, newCur *share.Node
	for _, n := range arr {
		if newHead == nil {
			newHead = n
			newCur = newHead
			continue
		}
		newCur.Next = n
		newCur = n
	}
	newCur.Next = nil
	tail := newCur
	newList := &share.SigleLinkedList{
		Head: newHead,
		Tail: tail,
	}

	share.DebugPrint("sorted list:", newList)
	return newList
}

// 按中轴将小者放左边，大者放右边
// 第一种方法，用了两个外层的 for
//     第一次排了 < pivot 和 >= pivot
//     第二次排了 == pivot 和 > pivot
// 返回 >= pivot 的起始位置，和 > pivot 的起始位置
func splitArrByPivot_1(arr []*share.Node, pivot int) (int, int) {
	// 类似快排的分段
	i := 0
	j := len(arr) - 1
	gteStart := 0
	for i < j {
		for arr[i].Value.(int) < pivot && i < j {
			i++
			gteStart++
		}
		for arr[j].Value.(int) >= pivot && i < j {
			j--
		}

		if i < j { // swap i, j
			swap(arr, i, j)
		}
	} // 前部分已排好
	share.DebugPrint("step, sorted 1st-part:", arr, gteStart)

	// 排后半部分
	i = gteStart
	j = len(arr) - 1
	gtStart := gteStart - 1
	for i < j {
		for arr[i].Value.(int) == pivot && i < j {
			i++
			gtStart++
		}

		for arr[j].Value.(int) > pivot && i < j {
			j--
		}

		if i < j {
			swap(arr, i, j)
		}
	}
	share.DebugPrint("step, ordered 2nd-part:", arr, gtStart)
	return gteStart, gtStart
}

func swap(arr []*share.Node, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// 按中轴将小者放左边，大者放右边
// 第二种方法，只进行一次循环
// 返回 >= pivot 的起始位置，和 > pivot 的起始位置
func splitArrByPivot_2(arr []*share.Node, pivot int) (int, int) {
	// 类似快排的分段
	i := 0
	ltEnd := -1
	pivotEnd := -1
	gtStart := len(arr)
	// i 左边全是 <= pivot 的值
	// 且 < 的全在 = 的左边
	// i 右边全是 > pivot 的值
	for i < gtStart {
		for arr[i].Value.(int) < pivot && i < gtStart { // 每一个判断上都要带 i<gtStart
			swap(arr, i, ltEnd+1)
			i++
			ltEnd++
		}
		fmt.Println("i, ltend", i, ltEnd, arr)
		for arr[i].Value.(int) == pivot && i < gtStart { // 因为每一个 for 都有可能更新 i 或 gtStart
			i++
			pivotEnd++
		}
		for arr[i].Value.(int) > pivot && i < gtStart { // 若不然，就需要每一步改动都对外层 continue
			swap(arr, i, gtStart-1)
			gtStart--
		}
	}
	share.DebugPrint("step, sorted:", arr)

	gteStart := gtStart
	if ltEnd < pivotEnd {
		gteStart = ltEnd + 1
	}
	return gteStart, gtStart
}
