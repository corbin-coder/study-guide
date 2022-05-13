package basics

import "fmt"

func pointer() {
	data := 10
	ptr := &data
	fmt.Println("Value stored in variable is", data)  //10
	fmt.Println("Value stored in variable is", *ptr)  //10
	fmt.Println("Value stored in variable is", &data) //0xc04200e210
	fmt.Println("Value stored in variable is", ptr)   //0xc04200e210

}

func passByRef() {
	i := 10
	fmt.Println("value before ref increment: ", i) //10
	incrementByVal(i)
	fmt.Println("value before ref increment: ", i) //10

	fmt.Println("value before ref increment: ", i) //10
	incrementByRef(&i)
	fmt.Println("value after ref increment: ", i) //11

}

func incrementByVal(x int) {
	x++
}

func incrementByRef(ptr *int) {
	(*ptr)++
}
