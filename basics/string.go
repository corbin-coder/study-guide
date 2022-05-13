package basics

import "fmt"

func updateString() {
	s := "helloWorld"
	r := []rune(s)
	r[0] = 'H'
	s = string(r)
	fmt.Print(s)
}
