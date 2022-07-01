package examples

import (
	"fmt"
	"sort"
)

func getMajority(data []int) (int, bool) {
	size := len(data)
	majndex := len(data)
	sort.Ints(data)
	candidate := data[majndex]
	count := 0

	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return data[majndex], true
	}
	fmt.Println("majority does not exist")
	return 0, false

}
