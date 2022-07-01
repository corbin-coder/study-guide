package sorting

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	BubbleSort(data)
	t.Log(data)
}
