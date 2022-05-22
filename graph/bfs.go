package graph

import "github.com/corbin-coder/codingtest/queue"

func (g *Graph) BFSQueue(index int, visited []int) {
	var curr int
	que := new(queue.Queue)
	visited[index] = 1

	que.Add(index)

	for que.Size() != 0 {
		curr, _ = que.Remove()
		head := g.VartexList[curr].head
		for head != nil {
			if visited[head.destination] == 0 {
				visited[head.destination] = 1
				que.Add(head.destination)
			}
			head = head.next
		}
	}

}

func (g *Graph) BFS() {
	count := g.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}

	for i := 0; i < count; i++ {
		if visited[i] == 0 {
			g.BFSQueue(i, visited)
		}
	}
}
