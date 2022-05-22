package sorting

//not a stable algorithm
//when data is random
func swap(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}

func quickSortUtil(arr []int, lower, upper int, comp func(int, int) bool) {
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper

	for lower < upper {
		for comp(arr[lower], pivot) == false && lower < upper {
			lower++
		}
		for comp(arr[upper], pivot) && lower <= upper {
			upper--
		}
		if lower < upper {
			swap(arr, upper, lower)
		}
		swap(arr, upper, start)
		quickSortUtil(arr, start, upper-1, comp)
		quickSortUtil(arr, upper+1, stop, comp)
	}
}

func quickSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	quickSortUtil(arr, 0, size-1, comp)
}
