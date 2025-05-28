package functionexamples

import "fmt"

func MultipleReturns() (int, int) {
	fmt.Println("multiple returns are possible, but must be declared in brackets")
	return 1, 2
}

func ReturnNamesVals() (x, y, z int) {
	fmt.Println("outputs can be named and declared in the code instead of returning them manually")
	x, y, z = 1, 3, 5
	return
}

func TypesOfCalls() {
	fmt.Println("functions args can be values or references")
	num1, num2 := 2, 5

	res1 := multiplyValues(num1, num2)
	fmt.Println(res1, num1, num2)

	fmt.Println("use & before variable name when passing by reference")
	fmt.Println(multiplyOneReference(&num1, num2), num1, num2)

	fmt.Println(multiplyTwoReferences(&num1, &num2), num1, num2)
}

func VariadicFunction() {
	fmt.Println(`can only have one variadic param per function
	AND it must be the last param in the list.
	number of variadic params can be zero`)

	variadicExample("4 nums", 1, 2, 3, 4)
	variadicExample("1 num", 1)
	variadicExample("no nums")
}

func variadicExample(title string, nums ...int) {
	fmt.Print(title)

	for _, n := range nums {
		fmt.Print(" Number: ", n)
	}

	fmt.Println("")
}

func multiplyValues(a, b int) int {
	fmt.Println("changing a in here doesn't change num1")
	a *= 2
	return a * b
}

func multiplyOneReference(a *int, b int) int {
	fmt.Println("changing a in here changes num1, but b doesn't change num2")
	*a *= 2
	b *= 2
	return *a * b
}

func multiplyTwoReferences(a, b *int) int {
	fmt.Println("changing a and b here WILL change num1 and num2")
	*a *= 2
	*b *= 2
	return *a * *b
}
