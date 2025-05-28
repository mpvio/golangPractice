package functionexamples

import "fmt"

func UsingAnonyFuncs() {
	funcFromFunction := ReturnsAnonyFunc()
	passedFuncAsParam := TakeAnonAsInput(funcFromFunction, 5)
	fmt.Println(passedFuncAsParam)
}

func ReturnsAnonyFunc() func(x int) int {
	value := func(x int) int {
		return x * x
	}

	// this function is called with value()
	fmt.Println("example of value:", value(2))

	//this is the equivalent of defining value2,
	// then calling value2(10)
	value2 := func(y int) int {
		return y + 10
	}(10)

	fmt.Println("example of anon func with input provided:", value2)
	return value
}

func TakeAnonAsInput(funcName func(x int) int, input int) int {
	fmt.Println("funcName is a function taken as input that takes one int as param")
	// execute the function taken as param on the other input taken as param
	return funcName(input)
}
