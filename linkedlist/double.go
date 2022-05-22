package linkedlist

import "fmt"

type DoublyLinkedList struct {
	head  *DNode
	tail  *DNode
	count int
}

type DNode struct {
	value int
	next  *DNode
	prev  *DNode
}

func (list *DoublyLinkedList) Size() int {
	return list.count
}

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.count == 0
}

func (list *DoublyLinkedList) Peek() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	return list.head.value, true
}

func (list *DoublyLinkedList) AddHead(value int) {
	newNode := &DNode{value, nil, nil}
	if list.count == 0 {
		list.tail = newNode
		list.head = newNode
	} else {
		list.head.prev = newNode
		newNode.next = list.head
		list.head = newNode
	}
	list.count++
}

func (list *DoublyLinkedList) AddTail(value int) {
	newNode := &DNode{value, nil, nil}
	if list.count == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
	list.count++
}

func (list *DoublyLinkedList) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("emptylisterror")
		return 0, false
	}
	value := list.head.value
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}
	list.count--
	return value, true
}

func (list *DoublyLinkedList) DeleteNode(key int) bool {
	curr := list.head
	if curr == nil {
		return false
	}
	if curr.value == key {
		curr = curr.next
		list.count--
		if curr != nil {
			list.head = curr
			list.head.prev = nil
		} else {
			list.tail = nil //only one element in list
		}
	}
	for curr.next != nil {
		if curr.next.value == key {
			curr.next = curr.next.next
			if curr.next == nil { //last element case
				list.tail = curr
			} else {
				curr.next.prev = curr
			}
			list.count--
			return true
		}
		curr = curr.next
	}
	return false
}

func (list *DoublyLinkedList) IsPresent(key int) bool {
	temp := list.head
	for temp != nil {
		if temp.value == key {
			return true
		}
		temp = temp.next
	}
	return false
}

func (list *DoublyLinkedList) FreeList() {
	list.tail = nil
	list.head = nil
	list.count = 0
}

func (list *DoublyLinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

func (list *DoublyLinkedList) CopyListReversed(dll *DoublyLinkedList) {
	curr := list.head
	for curr != nil {
		dll.AddHead(curr.value)
		curr = curr.next
	}
}

func (list *DoublyLinkedList) CopyList(dll *DoublyLinkedList) {
	curr := list.head
	for curr != nil {
		dll.AddTail(curr.value)
		curr = curr.next
	}
}

func (list *DoublyLinkedList) SortedInsert(value int) {
	temp := &DNode{}
	curr := list.head
	if curr == nil {
		list.head = temp
		list.tail = temp
	}

	if list.head.value <= value { //at begining
		temp.next = list.head
		list.head.prev = temp
		list.head = temp
	}

	for curr.next != nil && curr.next.value > value { //treversal
		curr = curr.next
	}

	if curr.next == nil { //at the end
		list.tail = temp
		temp.prev = curr
		curr.next = temp
	} else { //all other
		temp.next = curr.next
		temp.prev = curr
		curr.next = temp
		temp.next.prev = temp
	}
}

func (list *DoublyLinkedList) RemoveDuplicate() {
	curr := list.head
	var deleteMe *DNode
	for curr != nil {
		if (curr.next != nil) && curr.value == curr.next.value {
			deleteMe = curr.next
			curr.next = deleteMe.next
			curr.next.prev = curr
			if deleteMe == list.tail {
				list.tail = curr
			}
		} else {
			curr = curr.next
		}
	}
}
