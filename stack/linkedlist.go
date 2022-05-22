package stack

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(value int) {
	s.head = &Node{value, s.head}
	s.size++
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyError")
		return 0, false
	}
	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}

func (s *Stack) Print() {
	temp := s.head
	fmt.Print("value stored in stack are :: ")
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}
