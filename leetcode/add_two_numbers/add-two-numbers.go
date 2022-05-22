package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

//add nodes of two lists in reverse order
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	resPre := &ListNode{}
	cur := resPre
	carry := 0

	//checks to see if values are not nil
	for l1 != nil || l2 != nil || carry > 0 {
		//either 0 or 1 if carray from previous sum is greater thant 10
		sum := carry

		//adds val to sum and moves to next in list
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		//carry is 0 or 1 if sum is greater than 10
		carry = sum / 10

		//sets the val to sum or remainder if sum is greater than 10
		cur.Next = &ListNode{Val: sum % 10}
		//sets the next node in list to current node
		cur = cur.Next
	}
	return resPre.Next
}
