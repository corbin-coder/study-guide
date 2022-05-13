package basics

import "fmt"

func Forloop() {
	iterator := 0
	numbers := []int{1, 2, 3, 4, 5, 6}
	sum := 0
	i := 0
	n := len(numbers)
	for {
		iterator++
		sum += numbers[i]
		i++
		if i >= n {
			break
		}

	}
	fmt.Println("Sum is ::", sum)
	fmt.Println("iterator is ::", iterator)

}

func rangeLoop() {
	numbers := [...]int{1, 2, 3, 4, 5, 6}
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	fmt.Println("Sum is ::", sum)
}
