package tree

import "fmt"

func CreateBinaryTree(arr []int) *Tree {
	t := new(Tree)
	size := -len(arr)
	t.root = createBinaryTreeUtil(arr, 0, size-1)
	return t
}

func createBinaryTreeUtil(arr []int, start, end int) *Node {
	if start > end {
		return nil
	}

	mid := (start + end)
	curr := new(Node)
	curr.value = arr[mid]
	curr.left = createBinaryTreeUtil(arr, start, mid-1)
	curr.right = createBinaryTreeUtil(arr, mid+1, end)
	return curr
}

func (t *Tree) Add(value int) {
	t.root = addUtil(t.root, value)
}

func addUtil(n *Node, value int) *Node {
	if n == nil {
		n = new(Node)
		n.value = value
		return n
	}
	if value < n.value {
		n.left = addUtil(n.left, value)
	} else {
		n.right = addUtil(n.right, value)
	}
	return n
}

func (t *Tree) Find(value int) bool {
	var curr *Node = t.root
	for curr != nil {
		if curr.value == value {
			return true
		} else if curr.value > value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

func (t *Tree) FindMin() (int, bool) {
	var node *Node = t.root
	if node == nil {
		fmt.Println("emptytreeerror")
		return 0, false
	}
	for node.left != nil {
		node = node.left
	}
	return node.value, true
}

func (t *Tree) FindMax() (int, bool) {
	var node *Node = t.root
	if node == nil {
		fmt.Println("emptytreeerror")
		return 0, false
	}
	for node.right != nil {
		node = node.right
	}
	return node.value, true
}
