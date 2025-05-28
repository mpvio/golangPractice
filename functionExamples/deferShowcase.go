package functionexamples

import "fmt"

func DeferenceShowcase() {
	fmt.Println(func2())
	//defer -> run this after rest of function
	defer fmt.Println(func2())
	//defers are LIFO, so this will run before above
	defer fmt.Println(func1())
}

func func1() int {
	return 10
}

func func2() string {
	return "ten"
}
