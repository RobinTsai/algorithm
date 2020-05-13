# algorithm

## 链表问题（linkedList）

- 对于笔试，不用太在乎空间复杂度，一切为了时间复杂度
- 对于面试，时间复杂度放在第一位，但一定要找到空间最省的方法

这是因为笔试时，是无法估量你的空间复杂度的。而面试时，自己可以找到最优的空间复杂度将是一个加分项。

链表面试常用的数据结构和技巧：

- 使用容器（哈希表、数组等）
- 快慢指针

### 面试题

#### 找中点问题

- 输入链表头节点，奇数长度返回中点，偶数长度返回上中点（[findMidValue](./linkedList/findMidValue/main.go#getMidOrFirstMid)）
- 输入链表头节点，奇数长度返回中点，偶数长度返回下中点（[findMidValue](./linkedList/findMidValue/main.go#getMidOrSecondMid)）
- 输入链表头节点，奇数长度返回中点前一个节点，偶数长度返回上中点前一个节点（[findMidValue](./linkedList/findMidValue/main.go#getPrevMidOrPrevFirstMid)）
- 输入链表头节点，奇数长度返回中点前一个节点，偶数长度返回下中点后一个节点（[findMidValue](./linkedList/findMidValue/main.go#getPrevMidOrNextSecondMid)）

找中点问题用快慢指针来做，以第一个题为例，一个快指针每次走两步，一个慢指针每次走一步，当快指针走到尽头时，慢指针指向的就是中点位置。对于奇数长度或偶数长度属于细节问题。

#### 判断链表暖是否回文

- 给定一个单链表头节点 head，请判断该链表是否为回文结构
    - 栈的方法（笔试用）
    - 改原链表的方法（面试用）

方法一：用栈结构，将链表逐个 push 进栈，然后再次遍历，同时从栈中 pop 出来，逐个检查是否相等。（[isPalindrome](./linkedList/isPalindrome/main.go#isPalindrome_1)）

方法二：比方法一节省使用栈空间，使用快慢指针。快指针每次走 2 步，慢指针每次走 1 步并压栈，当快指针走完时用栈弹出值与慢指针的每个做比较。（[isPalindrome](./linkedList/isPalindrome/main.go#isPalindrome_2)）

方法三：改原链表，空间复杂度为 O(1)。如方法二找到中点后，将后半部分的链表反向，最后再对两个列表逐个比较。但注意，此方法会更改原链表，默认情况下需要最后将链表再还原回去。（[isPalindrome](./linkedList/isPalindrome/main.go#isPalindrome_3)）