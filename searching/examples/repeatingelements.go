package examples

import "fmt"

func printRepeating(data []int) {
	s := make(map[int]struct{})

	for i := 0; i < len(data); i++ {
		_, ok := s[data[i]]
		if ok == true {
			fmt.Print(" ", data[i])
		} else {
			s[data[i]] = struct{}{}
		}
	}

}
