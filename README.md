# algorithm

## 数组问题（array）

### 双指针法

双指针法，即用快慢两个指针进行遍历操作，一般快指针在前面“探路”，慢指针会慢几拍向前移动。一般复杂度在 O(n)。

- 双指针法求数组某几个元素和为某个值（[fourSum](./array/doublePointer/fourSum.go)）

## 链表问题（linkedList）

- 对于笔试，不用太在乎空间复杂度，一切为了时间复杂度
- 对于面试，时间复杂度放在第一位，但一定要找到空间最省的方法

这是因为笔试时，是无法估量你的空间复杂度的。而面试时，自己可以找到最优的空间复杂度将是一个加分项。

链表面试常用的数据结构和技巧：

- 使用容器（哈希表、数组等）
- 快慢指针

### 面试题

#### 找中点问题

- 输入链表头节点，奇数长度返回中点，偶数长度返回上中点（[findMidValue/getMidOrFirstMid](./linkedList/findMidValue/main.go)）
- 输入链表头节点，奇数长度返回中点，偶数长度返回下中点（[findMidValue/getMidOrSecondMid](./linkedList/findMidValue/main.go)）
- 输入链表头节点，奇数长度返回中点前一个节点，偶数长度返回上中点前一个节点（[findMidValue/getPrevMidOrPrevFirstMid](./linkedList/findMidValue/main.go)）
- 输入链表头节点，奇数长度返回中点前一个节点，偶数长度返回下中点后一个节点（[findMidValue/getPrevMidOrNextSecondMid](./linkedList/findMidValue/main.go)）

找中点问题用快慢指针来做，以第一个题为例，一个快指针每次走两步，一个慢指针每次走一步，当快指针走到尽头时，慢指针指向的就是中点位置。对于奇数长度或偶数长度属于细节问题。

#### 判断链表是否回文

- 给定一个单链表头节点 head，请判断该链表是否为回文结构
    - 栈的方法（笔试用）
    - 改原链表的方法（面试用）

方法一：用栈结构，将链表逐个 push 进栈，然后再次遍历，同时从栈中 pop 出来，逐个检查是否相等。（[isPalindrome/isPalindrome_1](./linkedList/isPalindrome/main.go)）

方法二：比方法一节省使用栈空间，使用快慢指针。快指针每次走 2 步，慢指针每次走 1 步并压栈，当快指针走完时用栈弹出值与慢指针的每个做比较。（[isPalindrome/isPalindrome_2](./linkedList/isPalindrome/main.go)）

方法三：改原链表，空间复杂度为 O(1)。如方法二找到中点后，将后半部分的链表反向，最后再对两个列表逐个比较。但注意，此方法会更改原链表，默认情况下需要最后将链表再还原回去。（[isPalindrome/isPalindrome_3](./linkedList/isPalindrome/main.go)）

#### 链表分段

- 将单链表按某值划分成左边小、中间相等、右边大的形式
    - 将链表放入数组中，在数组上划分（笔试用）
    - 分成小、中、在三部分，再将三部分串起来（面试用）

方法一：用数组结构做中转，按中轴拆分为三段，再拼接起来。（[shardByValue_1](./linkedList/shardByValue/main.go)）


方法二：在原链表上进行操作（TODO）

> 方法一内部提供了两种将数组分段的方法，可参考 [splitArrByPivot_1](./linkedList/shardByValue/main.go/splitArrByPivot_1)、[splitArrByPivot_2](./linkedList/shardByValue/main.go/splitArrByPivot_2)

## 树问题（tree）

### 概念

- 二叉树，每个节点最多有两个子节点的树
- 满二叉树，树的每一层都是满的（不能再加一个节点）
- 完全二叉树，二叉树除最后一层外都是满的，且最后一层从左到右是依次有节点的（不一定满）
- 先序、中序、后序参考下图

![tree_three_orders](https://upload-images.jianshu.io/upload_images/3491218-8de0fc68fe963f46.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 基本操作

- 二叉树的先序、中序、后序遍历
    - 原生的方式 [traverse/raw](./tree/traverse/raw.go)
    - 递归用队列和不用队列的方式 [traverse/PreOrder](./tree/traverse/recursion.go)
- 按层遍历（宽度优先遍历）
    - 用普通数组的方式 [traverse/TraverseByLayer_ConvertToEdges](tree/traverse/layer.go)
    - 用队列 [traverse/TraverseByLayer_Queue](tree/traverse/layer.go)
- 序列化和反序列化，这里是转换成数组 [traverse/serialization](tree/traverse/serialization.go)
- 打印一个二叉树，按树形输出在终端上 [traverse/PrintAsTree](tree/traverse/printTree.go)
- 高度优先遍历（用栈）
- 二叉树的递归套路问题（从左右子树取信息）
    - 判断二叉树是否为平衡树 [isBalance](tree/isBalance/main.go)
    - 获取二叉树某一节点的所有后继节点 [successor](tree/successor/getSuccessors.go)
    - 二叉树的最长路径 [theLongestDistance](tree/theLongestDistance/main.go)

附： 实现了一个简单的队列 [Queue](./tree/share/simpleQueue.go)

## 图问题

- 深度优先遍历（DFS，Depth-First-Search），遍历顺序是按层向下遍历的（队列）；
- 广度优先遍历（BFS，Breadth-First-Search），遍历顺序是先某一路走到底，再走下一路进行遍历（栈）。
