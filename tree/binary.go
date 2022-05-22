package tree

import (
	"fmt"
	"math"

	"github.com/corbin-coder/codingtest/stack"
)

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

func levelOrderBinaryTree(arr []int, start, size int) *Node {
	curr := &Node{arr[start], nil, nil}

	left := 2*start + 1
	right := 2*start + 2

	if left < size {
		curr.left = levelOrderBinaryTree(arr, left, right)
	}
	if right < size {
		curr.right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}

func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.value, " ")
	printPreOrder(n.left)
	printPreOrder(n.right)
}

func (t *Tree) PrintPreOrder() {
	printPreOrder(t.root)
	fmt.Println()
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	}
	printPostOrder(n.left)
	printPostOrder(n.right)
	fmt.Print(n.value)
}

func (t *Tree) PrintPostOrder() {
	printPostOrder(t.root)
	fmt.Println()
}

func printInOrder(n *Node) {
	if n == nil {
		return
	}
	printInOrder(n.left)
	fmt.Print(n.value)
	printInOrder(n.right)
}

func (t *Tree) PrintInOrder() {
	printInOrder(t.root)
	fmt.Print()
}

func treeDepth(root *Node) int {
	if root == nil {
		return 0
	}
	lDepth := treeDepth(root.left)
	rDepth := treeDepth(root.right)

	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

func nthPreOrder(node *Node, index int, counter *int) {
	if node != nil {
		(*counter)++
		if *counter == index {
			fmt.Println(node.value)
		}
		nthPreOrder(node.left, index, counter)
		nthPreOrder(node.right, index, counter)
	}
}

func (t *Tree) nthPreOrder(index int) {
	var counter int
	nthPreOrder(t.root, index, &counter)
}

func nthPostOrder(node *Node, index int, counter *int) {
	if node != nil {
		nthPostOrder(node.left, index, counter)
		nthPostOrder(node.right, index, counter)
		(*counter)++
		if *counter == index {
			fmt.Println(node.value)
		}
	}
}

func (t *Tree) NthPostOrder(index int) {
	var counter int
	nthPostOrder(t.root, index, &counter)
}

func nthInOrder(node *Node, index int, counter *int) {
	if node != nil {
		nthInOrder(node.left, index, counter)
		*counter++
		if *counter == index {
			fmt.Println(node.value)
		}
		nthInOrder(node.right, index, counter)
	}
}

func (t *Tree) NthInOrder(index int) {
	var counter int
	nthInOrder(t.root, index, &counter)

}

func copyTree(curr *Node) *Node {
	var temp *Node
	if curr != nil {
		temp = new(Node)
		temp.value = curr.value
		temp.left = copyTree(curr.left)
		temp.right = copyTree(curr.right)
		return temp
	}
	return nil
}

func (t *Tree) CopyTree() *Tree {
	Tree2 := new(Tree)
	Tree2.root = copyTree(t.root)
	return Tree2
}

func copyMirrorTree(curr *Node) *Node {
	var temp *Node
	if curr != nil {
		temp = new(Node)
		temp.value = curr.value
		temp.right = copyMirrorTree(curr.left)
		temp.left = copyMirrorTree(curr.right)
		return temp
	}
	return nil
}

func (t *Tree) CopyMirrorTree() *Tree {
	tree := new(Tree)
	tree.root = copyMirrorTree(t.root)
	return tree
}

func numNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	return (1 + numNodes(curr.right) + numNodes(curr.left))
}

func (t *Tree) NumNodes() int {
	return numNodes(t.root)
}

func numLeafNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	if curr.left == nil && curr.right == nil {
		return 1
	}
	return (numLeafNodes(curr.right) + numLeafNodes(curr.left))
}

func (t *Tree) NumLeafNodes() int {
	return numLeafNodes(t.root)
}

func isEqual(node1, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	} else {
		return ((node1.value == node2.value) && isEqual(node1.left, node2.left) && isEqual(node1.right, node2.right))
	}
}

func (t *Tree) IsEqual(t2 *Tree) bool {
	return isEqual(t.root, t2.root)
}

func printAllPath(curr *Node, stk *stack.Stack) {
	if curr == nil {
		return
	}
	stk.Push(curr.value)
	if curr.left == nil && curr.right == nil {
		stk.Print()
		stk.Pop()
		return
	}
	printAllPath(curr.right, stk)
	printAllPath(curr.left, stk)
	stk.Pop()
}

func findMaxBT(curr *Node) int {
	if curr == nil {
		return math.MinInt32
	}
	max := curr.value
	left := findMaxBT(curr.left)
	right := findMaxBT(curr.right)
	if left > max {
		max = left
	}
	if right > max {
		max = right
	}
	return max
}

func SearchBT(root *Node, value int) bool {
	var left, right bool
	if root == nil || root.value == value {
		return false
	}
	left = SearchBT(root.left, value)
	if left {
		return true
	}
	right = SearchBT(root.right, value)
	if right {
		return true
	}
	return false
}

func nummFullNodesBT(curr *Node) int {
	var count int
	if curr == nil {
		return 0
	}
	count = nummFullNodesBT(curr.right) + nummFullNodesBT(curr.left)
	if curr.right != nil && curr.left != nil {
		count++
	}
	return count
}

func summAllBT(curr *Node) int {
	var sum, leftSum, rightSum int
	if curr == nil {
		return 0
	}

	rightSum = summAllBT(curr.right)
	leftSum = summAllBT(curr.left)
	sum = rightSum + leftSum + curr.value
	return sum
}

func InitTree() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	LevelOrderBinaryTree(arr)
}
