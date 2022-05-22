package linkedlist

import "fmt"

type CircularLinkedList struct {
	tail  *CNode
	count int
}

type CNode struct {
	value int
	next  *CNode
}

func (list *CircularLinkedList) size() int {
	return list.count
}

func (list *CircularLinkedList) IsEmpty() bool {
	return list.count == 0
}

func (list *CircularLinkedList) Peek() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("emptylisterror")
		return 0, false
	}
	return list.tail.next.value, true
}

func (list *CircularLinkedList) AddHead(value int) {
	temp := &CNode{value, nil}
	if list.IsEmpty() {
		list.tail = temp
		temp.next = temp
	} else {
		temp.next = list.tail.next
		list.tail.next = temp
	}
	list.count++
}

func (list *CircularLinkedList) AddTail(value int) {
	temp := &CNode{value, nil}
	if list.IsEmpty() {
		list.tail = temp
		temp.next = temp
	} else {
		temp.next = list.tail.next
		list.tail.next = temp
		list.tail = temp
	}
	list.count++
}

func (list *CircularLinkedList) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("emptylisterror")
		return 0, false
	}
	value := list.tail.next.value
	if list.tail == list.tail.next {
		list.tail = nil
	} else {
		list.tail.next = list.tail.next.next
	}
	list.count--
	return value, true

}

func (list *CircularLinkedList) IsPresent(data int) bool {
	temp := list.tail
	for i := 0; i < list.count; i++ {
		if temp.value == data {
			return true
		}
		temp = temp.next
	}
	return false
}

func (list *CircularLinkedList) Print() {
	if list.IsEmpty() {
		return
	}
	temp := list.tail.next
	for temp != list.tail {
		fmt.Println(temp.value, " ")
		temp = temp.next
	}
	fmt.Println(temp.value)
}

func (list *CircularLinkedList) FreeList() {
	list.tail = nil
	list.count = 0
}

func (list *CircularLinkedList) RemoveNode(key int) bool {
	if list.IsEmpty() {
		return false
	}
	prev := list.tail
	curr := list.tail.next
	head := list.tail.next

	if curr.value == key { //head and single node case
		if curr == curr.next { //single node case
			list.tail = nil
		} else { //head case
			list.tail.next = list.tail.next.next
		}
		return true
	}

	prev = curr
	curr = curr.next

	for curr != head {
		if curr.value == key {
			if curr == list.tail {
				list.tail = prev
			}
			prev.next = curr.next
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false

}
