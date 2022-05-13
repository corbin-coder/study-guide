package basics

import "fmt"

func initSwitch() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("something else")
	}
}
