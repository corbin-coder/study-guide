package linkedlist

import "fmt"

type List struct {
	head  *Node
	count int
}

type Node struct {
	value int
	next  *Node
}

func (list *List) Size() int {
	return list.count
}

func (list *List) IsEmpty() bool {
	return (list.count == 0)
}

func (list *List) AddHead(value int) {
	list.head = &Node{value: value, next: list.head}
	list.count++
}

func (list *List) AddTail(value int) {
	curr := list.head
	newNode := &Node{value: value, next: nil}

	if curr == nil {
		list.head = newNode
		return
	}

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (list *List) Print() {
	temp := list.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

func (list *List) SortedInsert(value int) {
	newNode := &Node{value, nil}
	curr := list.head

	if curr == nil || curr.value > value {
		newNode.next = list.head
		list.head = newNode
		return
	}

	for curr.next != nil && curr.next.value < value {
		curr = curr.next
	}

	newNode.next = curr.next
	curr.next = newNode
}

func (list *List) IsPresent(data int) bool {
	temp := list.head
	for temp != nil {
		if temp.value == data {
			return true
		}
		temp = temp.next
	}
	return false
}

func (list *List) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	value := list.head.value
	list.head = list.head.next
	list.count--
	return value, true
}

func (list *List) DeleteNode(delValue int) bool {
	temp := list.head
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return false
	}

	if delValue == list.head.value {
		list.head = list.head.next
		list.count--
		return true
	}
	for temp.next != nil {
		if temp.next.value == delValue {
			temp.next = temp.next.next
			list.count--
			return true
		}
		temp = temp.next
	}
	return false
}

func (list *List) DeleteNodes(delValue int) {
	currNode := list.head
	for currNode != nil && currNode.value == delValue {
		list.head = currNode.next
		currNode = list.head
	}
	for currNode != nil {
		nextNode := currNode.next
		if nextNode != nil && nextNode.value == delValue {
			currNode.next = nextNode.next
		} else {
			currNode = nextNode
		}
	}
}

func (list *List) FreeList() {
	list.head = nil
	list.count = 0
}

func (list *List) Reverse() {
	curr := list.head
	var prev, next *Node

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	list.head = prev

}

func (list *List) RemoveDuplicate() {
	curr := list.head
	for curr != nil {
		if curr.next != nil && curr.value == curr.next.value {
			curr.next = curr.next.next
		} else {
			curr = curr.next
		}
	}
}

func (list *List) CopyListReversed() *List {
	var tempNode, tempNode2 *Node

	curr := list.head
	for curr != nil {
		tempNode2 = &Node{curr.value, tempNode}
		curr = curr.next
		tempNode = tempNode2
	}
	ll2 := new(List)
	ll2.head = tempNode
	return ll2

}

func (list *List) CopyList() *List {
	var headNode, tailNode, tempNode *Node
	curr := list.head
	if curr == nil {
		ll2 := new(List)
		ll2.head = nil
		return ll2
	}
	headNode = &Node{curr.value, nil}
	tailNode = curr.next
	for curr != nil {
		tempNode = &Node{curr.value, nil}
		tailNode.next = tempNode
		tailNode = tempNode
		curr = curr.next
	}
	ll2 := new(List)
	ll2.head = headNode
	return ll2
}

func (list *List) CompareList(ll *List) bool {
	return list.compareListUtil(list.head, ll.head)

}

func (list *List) compareListUtil(head1 *Node, head2 *Node) bool {
	if head1 == nil && head2 == nil {
		return true
	} else if (head1 == nil) || (head2 == nil) || (head1.value != head2.value) {
		return false
	} else {
		return list.compareListUtil(head1.next, head2.next)
	}
}

func (list *List) FindLength() int {
	curr := list.head
	count := 0
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

func (list *List) NthNodeFromBegining(index int) (int, bool) {
	count := 0
	curr := list.head
	for curr != nil && count < index-1 {
		count++
		curr = curr.next
	}
	return curr.value, true
}

func (list *List) NthNodeFromEnd(index int) (int, bool) {
	size := list.FindLength()
	startIndex := size - index + 1
	return list.NthNodeFromBegining(startIndex)
}
