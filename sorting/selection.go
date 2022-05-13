package sorting

//this is not a stable sort for time complexity
func selectionSort(arr []int) {
	size := len(arr)
	var i, j, max int
	for i = 0; i < size-1; i++ {
		max = 0
		for j = 1; j < size-1-i; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		arr[size-1-i], arr[max] = arr[max], arr[size-1-i]
	}
}
