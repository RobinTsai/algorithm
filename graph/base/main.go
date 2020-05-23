/**
 * 有向图
 * 定义、转换成统一定义、取最小单向遍历路径
 */
package main

import (
	"fmt"
)

func main() {
	matrix := getEdgeMatrixExample()
	fmt.Println("origin matrix is:", matrix)

	graph := ConvertEdgeMatrixToGraph(matrix)
	fmt.Println("\ngraph is:\n", graph)

	idxMap, n2Matrix := graph.ConvertToN2Matrix()
	fmt.Println("\nlabel to idx:", idxMap)
	fmt.Println("n2Matrix is:", n2Matrix)
}

type Graph struct {
	Points map[int]*Node // 编号=>点（这里编号就是点的值）
	Edges  []*Edge       // 边
}

type Edge struct {
	Weight int   // 权重值
	From   *Node // 入点
	To     *Node // 出点
}

type Node struct {
	Value    int     // 值
	InCount  int     // 入度，入口点的个数
	OutCount int     // 出度，出口点的个数
	Nexts    []*Node // 下一个点
	Edges    []*Edge // 即存入口边也存出口边
}

// 打印函数
func (g Graph) String() string {
	points := "points:"
	for idx, point := range g.Points {
		points += fmt.Sprintf("\n  idx: %d, value: %d, inCount: %d, outCount: %d", idx, point.Value, point.InCount, point.OutCount)
	}
	edges := "\nedges: "
	for _, edge := range g.Edges {
		edges += fmt.Sprintf("\n  weight: %d, from: %d, to: %d", edge.Weight, edge.From.Value, edge.To.Value)
	}
	return points + edges
}

// 从 graph.Points 上取 node
// 这样，
//   一来，按地址取数据，更新数据不会有错
//   二来，在这里进行初始化数据操作，直接使用不会有错
func (g *Graph) getNode(idx int) *Node {
	_, ok := g.Points[idx]
	if !ok {
		g.Points[idx] = &Node{
			Value: idx,
			Nexts: make([]*Node, 0),
			Edges: make([]*Edge, 0),
		}
	}
	return g.Points[idx]
}

// 转换为统一 Graph
// 将特定的 Edge 矩阵转换成 Graph
// Edge 矩阵，每一个元素是一个三元素数组
//   [weight, from, to]，weight 表示边权重， from/to 是方向
func ConvertEdgeMatrixToGraph(matrix [][3]int) Graph {
	graph := Graph{
		Points: map[int]*Node{},
		Edges:  []*Edge{},
	}
	for _, originEdge := range matrix {
		weight := originEdge[0]
		from := originEdge[1]
		to := originEdge[2]

		// fmt.Println(weight, from, to)
		fromNode := graph.getNode(from)
		toNode := graph.getNode(to)

		newEdge := &Edge{
			Weight: weight,
			From:   fromNode,
			To:     toNode,
		}
		fromNode.OutCount++
		toNode.InCount++
		fromNode.Edges = append(fromNode.Edges, newEdge)
		toNode.Edges = append(toNode.Edges, newEdge)
		graph.Edges = append(graph.Edges, newEdge)
	}
	return graph
}

// 取有向图最短遍历路径
func (g Graph) getMinPath() []Edge {
	return nil
}

// 转换为 N*N 的矩阵
// 二维矩阵，下标(i, j)表示从 i 到 j 路径（一步）的权重值
// 无路径的取 int 最大值
// 返回，
// 	   第一个参数是标号对应到矩阵后的下标
//     第二个参数是矩阵
func (g Graph) ConvertToN2Matrix() (map[int]int, [][]int) {
	// 下标索引 map, 点位标号 => matrix的下标索引
	idxMap := map[int]int{} // 由于 g.Points 点可能下标不是从 0 开始的，而是一个固定标号，所以做一个 map
	idx := 0
	for key, _ := range g.Points {
		idxMap[key] = idx
		idx++
	}

	// 初始化矩阵，对应点位为 int 最大值
	length := len(g.Points)
	matrix := make([][]int, length)
	intMax := 1<<63 - 1
	for i := 0; i < length; i++ {
		matrix[i] = make([]int, length)
		for j := 0; j < length; j++ {
			matrix[i][j] = intMax
		}
	}

	// 设置点位
	for _, edge := range g.Edges {
		// fmt.Println("edge is", edge.Weight, edge.From.Value, edge.To.Value)
		fromIdx := idxMap[edge.From.Value]
		toIdx := idxMap[edge.To.Value]
		weight := edge.Weight
		matrix[fromIdx][toIdx] = weight
	}

	return idxMap, matrix
}

// ------------------------------------------------------------------------------

/*
EdgeMatrix 示例：
[
	[weight, from, to],
	[weight, from, to],
	[weight, from, to],
	...
]
*/
func getEdgeMatrixExample() [][3]int {
	return [][3]int{
		[3]int{5, 1, 2}, // 表示从点 1 到点 2 需要 5 单位计量
		[3]int{3, 2, 3},
		[3]int{4, 1, 4},
		[3]int{8, 4, 3},
		[3]int{5, 2, 1},
		[3]int{7, 1, 3},
	}
}
