package graph

import (
	"fmt"

	"github.com/corbin-coder/codingtest/stack"
)

func (g *Graph) DFSStack() {
	count := g.count
	visited := make([]int, count)
	var curr int
	stk := new(stack.Stack)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}

	visited[1] = 1
	stk.Push(1)

	for stk.Size() != 0 {
		curr, _ = stk.Pop()
		head := g.VartexList[curr].head
		for head != nil {
			if visited[head.destination] == 0 {
				visited[head.destination] = 1
				stk.Push(head.destination)
			}
			head = head.next
		}
	}
}

func (g *Graph) DFSRec(index int, visited []int) {
	head := g.VartexList[index].head
	fmt.Println(index)
	for head != nil {
		if visited[head.destination] == 0 {
			visited[head.destination] = 1
			g.DFSRec(head.destination, visited)
		}
		head = head.next
	}
}

func (g *Graph) DFS() {
	count := g.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	for i := 0; i < count; i++ {
		if visited[i] == 0 {
			visited[i] = 1
			g.DFSRec(i, visited)
		}
	}
}
