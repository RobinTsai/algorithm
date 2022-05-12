package main

// IsMirror 判断是否为镜像树专用的 Receiver
// 镜像树即按 head 结点左右对称的树
type IsMirror struct {
	Head *Node // 树的顶结点
}

func getOneMirrorTree() *IsMirror {
	a := &Node{Value: 1}
	b := &Node{Value: 2}
	B := &Node{Value: 2}
	c := &Node{Value: 3}
	C := &Node{Value: 3}
	d := &Node{Value: 4}
	D := &Node{Value: 4}

	a.Left = b
	a.Right = B
	b.Left = c
	c.Left = d
	B.Right = C
	C.Right = D

	return &IsMirror{a}
}

// Exec 分别传入镜像结点，判断两结点的子结点是否互为镜像
func (obj *IsMirror) Exec(left, right *Node) bool {
	if left == right && left == nil { // 判断指针，不仅仅是值
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return obj.Exec(left.Left, right.Right) && obj.Exec(left.Right, right.Left)
}
