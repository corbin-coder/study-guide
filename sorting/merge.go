package sorting

//best for sorting linked list
//always 0(nlogn)
//slower than quick sort
//when we want to merge two sorted linked lists
//also used for external sorting
func merge(arr []int, tempArray []int, lowerIndex int, middleIndex int, upperIndex int, comp func(int, int) bool) {
	lowerStart := lowerIndex
	lowerStop := middleIndex
	upperStart := middleIndex + 1
	upperStop := upperIndex
	count := lowerIndex
	for lowerStart <= lowerStop && upperStart <= upperStop {
		if comp(arr[lowerIndex], arr[upperStart]) == false {
			tempArray[count] = arr[lowerStart]
			lowerStart++
		} else {
			tempArray[count] = arr[upperStart]
			upperStart++
		}
		count++
	}
	for lowerStart <= lowerStop {
		tempArray[count] = arr[lowerStart]
		count++
		lowerStart++
	}
	for upperStart <= upperStop {
		tempArray[count] = arr[upperStart]
		count++
		upperStart++
	}
	for i := lowerIndex; i <= upperIndex; i++ {
		arr[i] = tempArray[i]
	}
}

func mergeSrt(arr []int, tempArray []int, lowerIndex int, upperIndex int, comp func(int, int) bool) {
	if lowerIndex >= upperIndex {
		return
	}
	middleIndex := (lowerIndex + upperIndex) / 2
	mergeSrt(arr, tempArray, lowerIndex, middleIndex, comp)
	mergeSrt(arr, tempArray, middleIndex+1, upperIndex, comp)
	merge(arr, tempArray, lowerIndex, middleIndex, upperIndex, comp)
}

func mergeSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	tempArray := make([]int, size)
	mergeSrt(arr, tempArray, 0, size-1, comp)
}
