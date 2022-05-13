package searching

import (
	"fmt"
	"sort"
)

type input struct {
	data  []int
	value int
}

//Must search through entire list until you find it
//Time complexity 0(n)
//Space complexity 0(1)
func (in input) unsortedLinearSearch() bool {
	size := len(in.data)
	for i := 0; i < size; i++ {
		if in.data[i] == in.value {
			return true
		}
	}
	return false
}

//you can stop searching when the index of value of data[] becomes bigger than value
//Time complexity 0(n)
//Space complexity 0(1)
func (in input) sortedLinearSearch() bool {

	size := len(in.data)
	for i := 0; i < size; i++ {
		if in.data[i] == in.value {
			return true
		} else if in.data[i] > in.value {
			return false
		}
	}
	return false
}

//Time Complexity 0(n)
//Space Complexity 0(n)
func (in input) findDuplicate() {
	s := make(map[int]struct{})
	size := len(in.data)
	for i := 0; i < size; i++ {
		if findHash(s, in.data[i]) == true {
			fmt.Print(" ", in.data[i])

		} else {
			s[in.data[i]] = struct{}{}
		}
	}
	fmt.Println()
}

func findHash(i map[int]struct{}, val int) bool {
	_, exists := i[val]
	return exists
}

//0(n logn)
//0(1)
func (in *input) getMax() int {
	size := len(in.data)
	max := in.data[0]
	maxCount := 1
	curr := in.data[0]
	currCount := 1
	sort.Ints(in.data)

	for i := 1; i < size; i++ {
		if in.data[i] == in.data[i-1] {
			currCount++
		} else {
			currCount = 1
			curr = in.data[i]
		}

		if currCount > maxCount {
			maxCount = currCount
			max = curr
		}
	}
	return max
}
