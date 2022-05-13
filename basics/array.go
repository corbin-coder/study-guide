package basics

import "fmt"

func arr() {
	var arr [10]int
	fmt.Println(arr)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	count := len(arr)
	fmt.Println("length of array", count) //10
}

func slice() {
	var s []int
	for i := 1; i <= 17; i++ {
		s = append(s, i)
		//cap() can change
		//len() is that actual lengh
	}
}
