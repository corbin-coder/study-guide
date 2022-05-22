package graph

import (
	"fmt"

	"github.com/corbin-coder/codingtest/queue"
)

type Edge struct {
	source      int
	destination int
	cost        int
	next        *Edge
}

type EdgeList struct {
	head *Edge
}

type Graph struct {
	count      int
	VartexList [](*EdgeList)
}

func (g *Graph) Init(cnt int) {
	g.count = cnt
	g.VartexList = make([]*EdgeList, cnt)
	for i := 0; i < cnt; i++ {
		g.VartexList[i] = new(EdgeList)
		g.VartexList[i].head = nil
	}
}

func (g *Graph) AddEdge(source, destination, cost int) {
	edge := &Edge{source, destination, cost, g.VartexList[source].head}
	g.VartexList[source].head = edge
}

func (g *Graph) AddEdgeUnwighted(source, destination int) {
	g.AddEdge(source, destination, 1)
}

func (g *Graph) AddBiEdge(source, destination, cost int) {
	g.AddEdge(source, destination, cost)
	g.AddEdge(destination, source, cost)
}

func (g *Graph) AddBiEdgeUnweighted(source, destination int) {
	g.AddBiEdge(source, destination, 1)
}

func (g *Graph) Print() {
	for i := 0; i < g.count; i++ {
		ad := g.VartexList[i].head
		if ad != nil {
			fmt.Print("vertex", i, "i connected to : ")
			for ad != nil {
				fmt.Print("[", ad.destination, ad.cost, "]")
				ad = ad.next
			}
			fmt.Println()
		}
	}
}

func (g *Graph) PathExists(source, destination int) bool {
	count := g.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	visited[source] = 1
	g.DFSRec(source, visited)
	return visited[destination] != 0
}

func (g *Graph) IsConnected() bool {
	count := g.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	visited[0] = 1
	g.DFSRec(0, visited)
	for i := 0; i < count; i++ {
		if visited[i] == 0 {
			return false
		}
	}
	return true
}

func (g *Graph) ShortestPath(source int) {
	var curr int
	count := g.count
	distance := make([]int, count)
	path := make([]int, count)
	que := new(queue.LQueue)
	for i := 0; i < count; i++ {
		distance[i] = -1
	}

	que.Add(source)
	distance[source] = 0
	path[source] = source
	if que.Size() != 0 {
		curr, _ = que.Remove()
		head := g.VartexList[curr].head
		for head != nil {
			if distance[head.destination] == -1 {
				distance[head.destination] = distance[curr] + 1
				path[head.destination] = distance[curr] + 1
				path[head.destination] = curr
				que.Add(head.destination)
			}
			head = head.next
		}
	}
	for i := 0; i < count; i++ {
		fmt.Println(path[i], "to", i, "weight", distance[i])
	}
}
