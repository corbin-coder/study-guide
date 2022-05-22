package queue

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LQueue struct {
	head *Node
	tail *Node
	size int
}

func (q *LQueue) Size() int {
	return q.size
}

func (q *LQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LQueue) Print() {
	temp := q.head
	for temp != nil {
		fmt.Println(temp.value, " ")
		temp = temp.next
	}
}

func (q *LQueue) Add(value int) {
	temp := &Node{value, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++

}

func (q *LQueue) Remove() (int, bool) {
	if q.IsEmpty() {
		fmt.Println("queueemptyerror")
		return 0, false
	}
	value := q.head.value
	q.head = q.head.next
	q.size--
	return value, true
}
