package graph

import (
	"container/list"
	"fmt"
)

type Graph struct {
	v   int
	adj [][]int
}

func NewGraph(v int) *Graph {
	graph := &Graph{v: v}
	graph.adj = make([][]int, v)
	for i := range graph.adj {
		graph.adj[i] = make([]int, 0)
	}
	return graph
}

func (g *Graph) AddEdge(s, t int) {
	g.adj[s] = append(g.adj[s], t)
	g.adj[t] = append(g.adj[t], s) // 无向图需要回指向
}

func (g *Graph) bfs(s, t int) {

	// 优先处理边界
	if s == t {
		return
	}
	visited := make([]bool, g.v)

	queue := list.New()
	// 推送源顶点进入队列
	queue.PushBack(s)

	// 设置初始节点为起步节点
	visited[s] = true

	// 用于回朔路径
	prev := make([]int, g.v)
	for i := range prev {
		prev[i] = -1
	}

	for queue.Len() > 0 {

		// 队列处理
		// 获取第一个元素
		element := queue.Front()
		w := element.Value.(int)
		queue.Remove(element)

		for _, q := range g.adj[w] {
			// 假设第一个为0的顶点 得到的二维数组里面为该0顶点与其他顶点的关系建立映射
			// 判断顶点里面点数据是否为被搜索过
			if !visited[q] {
				// 如果没有被搜索，还需要记录下搜索路径
				// 第一次w为0 表示着0对映射关系的顶点 是经过这个路径的
				// 用于回溯打印
				prev[q] = w
				if q == t {
					g.printPath(prev, t)
					return
				}
				// 标记改顶点已经搜索过
				visited[q] = true
				// 将顶点里面连接的数据按序放入queue中
				queue.PushBack(q)
			}
		}
	}

}

func (g *Graph) printPath(prev []int, t int) {
	fmt.Println(prev)
	if prev[t] != -1 {
		// 回朔规则是：-1 0 1 2 3 0
		// g.printPath(prev, 3)
		// g.printPath(prev, 2)
		// g.printPath(prev, 1)
		// g.printPath(prev, 0)
		// 当=-1以后 就会从0开始递归打印
		g.printPath(prev, prev[t])
	}
	fmt.Print(t, " ")
}

//func main() {
//	graph := NewGraph(6) // 初始化图，节点数量为示例
//
//	graph.AddEdge(0, 1)
//	graph.AddEdge(1, 2)
//	graph.AddEdge(2, 3)
//	graph.AddEdge(3, 4)
//	graph.AddEdge(4, 5)
//	graph.AddEdge(5, 0)
//
//	// i是顶点 l是顶点里面的链接信息
//	for i, l := range graph.adj {
//		fmt.Println(i, l)
//	}
//
//	graph.bfs(0, 3)
//}
