package basics

func max(x, y int) int {
	var max int
	if x > y {
		max = x
	} else {
		max = y
	}
	return max
}
