package main

import "fmt"

//explicit type conversion between number types and the d/f/g format types
func TypeConversion() {
	var num1 int = 5

	//conversions
	var num2 float64 = float64(num1)
	var num3 int64 = int64(num1)

	//number types must match to be calculated together
	var sum = num2 + float64(num3) + 0.5

	// %d for ints, %f for floats, %g will remove the decimal from a float if it's x.0
	fmt.Printf("%d, %f, %d, %g, %g", num1, num2, num3, sum, sum-0.5)
}

//initializing variables, default values, joint declaration + assignments via :=
func VariableExamples() {
	//no such thing as uninitialized: this imm. gets default value ("" for strings)
	var noValueGiven string

	//types can be inferred without being declared explicitly
	var stringTypeIsInferred = "Hello"
	var intTypeIsInferred = 10
	var float64TypeisInferred = 5.5

	//multiple variables can be declared in one line
	//int, string, float64
	var x, y, z = 1, "2", 3.0

	//including a type means all variables here MUST have that type
	var a, b int = 1, 2

	// := can be used to declare a variable AND assign a value to it right away
	// no including type here, it's inferred.
	// this is equivalent to var c int = 3
	// := should only be used inside a func, not globally.
	c := 3

	// this is valid but only because at least one of these variables is new (d)
	// cannot use := to declare a variable that's already been declared!
	a, b, c, d := 5, 5, 5, 5

	// functions can return multiple values, these can also both be captured by variables
	e, f := numbers()

	fmt.Println("Strings:", noValueGiven, stringTypeIsInferred, y)
	fmt.Println("Floats:", float64TypeisInferred, z, "Sum:", float64TypeisInferred+z)
	fmt.Println("Ints:", intTypeIsInferred, x, a, b, c, d, "Sum:", intTypeIsInferred+x+a+b+c+d)
	fmt.Println("These came from the same function:", e, f)

	fmt.Println("Remember: Global variables (i.e. outside funcs) MUST use const or var, never just 'x = y' or 'x := y'!")
}

func numbers() (int, int) {
	return 1, 1
}
