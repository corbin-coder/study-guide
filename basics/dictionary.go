package basics

import "fmt"

func dictionary() {
	m := make(map[string]int)
	m["apple"] = 40
	m["bannana"] = 30
	for key, val := range m {
		fmt.Print("[", key, " -> ", val, "]")
	}
	delete(m, "apple")
	value, ok := m["apple"]
	fmt.Println("apple price:", value, "Present:", ok) //0 false
}
