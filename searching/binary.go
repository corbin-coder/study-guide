package searching

//divides list in half if middle val is = search val returns true
//if middle val is greater than search val discards all greater vals
//if middle val is less than search val discards all lesser vals
//Time Complexity 0(logn)
//Space Compexity 0(1)
func (in input) binarySearch() bool {
	size := len(in.data)
	low := 0
	high := size - 1
	mid := 0

	for low <= high {
		mid = low + (high-low)/2
		if in.data[mid] == in.value {
			return true
		} else if in.data[mid] < in.value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

//Time Complexity 0(logn)
//Space Complexity 0(logn)
func (in input) recursiveBinarySearch(low, high int) bool {
	if low > high {
		return false
	}
	mid := low + (high-low)/2
	if in.data[mid] == in.value {
		return true
	} else if in.data[mid] < in.value {
		return in.recursiveBinarySearch(mid+1, high)
	} else {
		return in.recursiveBinarySearch(low, mid+1)
	}
}
