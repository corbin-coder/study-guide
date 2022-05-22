package sorting

func less(val1, val2 int) bool {
	return val1 < val2
}

func more(val1 int, val2 int) bool {
	return val1 > val2
}

//best when input is nearly sorted
func bubbleSort(arr []int) []int {
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if more(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}
