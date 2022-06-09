# 技巧

## 巧用 go 库自带的搜索方法

sort.Search 系列方法

```go
package sort

// n 指示了在 [0, n) 的范围中查找
// f 是匹配方法，需要保证目标位置 i：
//      在 i 之前的值 f 函数都为 false
//      在 i 及之后的位置 f 函数都为 true
// 会通过 二分查找法，找到最小满足 f 函数的下标值
// 如果没有符合条件的，将返回 n 值
func Search(n int, f func(int) bool) int

// 调用 Search(len(a), func(i int) bool { return a[i] >= x })
func SearchInts(a []int, x int) int

// 还有其他变体，总归来说都是调用的 Search 方法，知道它就行
```