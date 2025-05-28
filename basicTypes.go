package main

import "fmt"

//basic data types: numbers, strings, booleans

//print complex numbers AND as real/ imag components
func NumberExample() {
	// Numbers
	/*
		integers: signed, unsigned, different sizes like in Rust
		floats: float32, float64; decimal, exponential and mix of the two (e.g. 13.16e12)
		complex: imaginary numbers: complex64 <- float32, complex128 <- float64
	*/

	var a complex64 = complex(6, 2)
	var b complex128 = complex(9, 3)
	fmt.Println(a)
	fmt.Println(real(b), imag(b))
}

//showcase boolean checking, then formatted string
func BooleanExample() {
	str1 := "Hello"
	str2 := "World"
	str3 := "Hello"

	result1 := str1 == str2
	result2 := str1 == str3

	fmt.Println("Result1 is", result1, ", Result2 is", result2)
	// %formatting explained later
	fmt.Printf("Result1 is %T, Result2 is %T.\n", result1, result2)
}

//string, then different format types
func StringAndStringFormattingExample() {
	str := "Hello World"

	// %d -> number
	fmt.Printf("Length: %d.\n", len(str))
	// %s -> string
	fmt.Printf("String: %s.\n", str)
	// %T -> type
	fmt.Printf("Type: %T.\n", str)
	// %formatting can be varied
	fmt.Printf("String: %s. Type: %T.\n", str, str)
}

// since all objs use the any interface,
// can use similar logic to get underlying types
func TypeAssertions() {
	var value any = "string"

	var str string = value.(string)
	fmt.Println("just using value.(type) will crash if type is wrong:", str)

	fmt.Println("handle additional boolean output for error handling:")
	num, valid := value.(int)
	if valid {
		fmt.Println("this is a number!", num)
	} else {
		fmt.Println("this isn't a number and nothing will panic!", num)
		fmt.Println("since num was invalid above, just returned default int value (0)")
	}
}

// similar to above, can use switch statements on types
func TypeSwitches() {
	var v any = "string"
	switch t := v.(type) {
	//can switch on multiple criteria per case!
	case int, float32:
		fmt.Println("int or float32", t)
	case string, bool:
		fmt.Println("string or bool", t)
	default:
		fmt.Println("unknown", t)
	}
}
