package genericExample

import "fmt"

func UseGenericFunction() {
	var a int = 10
	var b float32 = 10

	fmt.Println("add what types are valid to generic function in [here]")
	fmt.Println("define a new type there and use it in the function (e.g. n)")
	fmt.Println(genericSquaring(a), genericSquaring(b))

	fmt.Println("can define multiple types as interface and pass that in too:")
	fmt.Println(moreGenericSquaring(a))

	fmt.Println("remember, 'any' is also an interface so can be used as well:")
	myPrint("hello")
	myPrint(100)
}

func genericSquaring[n int | float32](number n) n {
	return number * number
}

type allNumbers interface {
	int | float32 | float64
}

func moreGenericSquaring[n allNumbers](number n) n {
	return number * number
}

func myPrint[x any](input x) {
	fmt.Println(input)
}
