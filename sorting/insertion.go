package sorting

//must be small input
//faster than quick and merge sort if input is small
func insertionSort(arr []int) {
	size := len(arr)
	var temp, i, j int
	for i = 1; i < size; i++ {
		temp = arr[i]
		for j = i; j > 0 && more(arr[j-1], temp); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}
