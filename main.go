package main

import (
	"log"

	"github.com/corbin-coder/codingtest/sorting"
)

func main() {
	nums := []int{1, 2, 5, 4, 3}
	res := sorting.BubbleSort(nums)
	log.Println(res)
}
