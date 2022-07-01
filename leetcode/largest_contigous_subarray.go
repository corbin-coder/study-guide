package leetcode

import "log"

func rewrite3(nums []int) []int {
	finalTemp := []int{}
	temp := []int{}
	//curr := 0

	for i := 0; i != len(nums); i++ {
		log.Println("Start")
		if i == 0 || nums[i] == nums[i-1]+1 {
			log.Println("in loop")
			temp = append(temp, nums[i])
			log.Println(temp)
		} else {
			temp = []int{}
			temp = append(temp, nums[i])

		}
		if len(temp) >= len(finalTemp) {

			finalTemp = temp
		}
	}
	return finalTemp
}
