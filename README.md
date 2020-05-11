# algorithm

## 链表问题（linkedList）

- 对于笔试，不用太在乎空间复杂度，一切为了时间复杂度
- 对于面试，时间复杂度放在第一位，但一定要找到空间最省的方法

这是因为笔试时，是无法估量你的空间复杂度的。而面试时，自己可以找到最优的空间复杂度将是一个加分项。

链表面试常用的数据结构和技巧：

- 使用容器（哈希表、数组等）
- 快慢指针

面试题

- 输入链表头节点，奇数长度返回中点，偶数长度返回上中点（[findMidValue](./linkedList/findMidValue/main.go#getMidOrFirstMid)）
- 输入链表头节点，奇数长度返回中点，偶数长度返回下中点（[findMidValue](./linkedList/findMidValue/main.go#getMidOrSecondMid)）
- 输入链表头节点，奇数长度返回中点前一个节点，偶数长度返回上中点前一个节点（[findMidValue](./linkedList/findMidValue/main.go#getPrevMidOrPrevFirstMid)）
- 输入链表头节点，奇数长度返回中点前一个节点，偶数长度返回下中点后一个节点（[findMidValue](./linkedList/findMidValue/main.go#getPrevMidOrNextSecondMid)）
