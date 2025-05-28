package arraysAndSlices

import "fmt"

func ArrayFuncs() {
	creatingArray()
	copyingArray()
}

func init() {
	fmt.Println("init run in arraysLesson. is first before main's inits and a < s of slicesLesson")
}

func creatingArray() {
	fmt.Println("1. array format: var array_name[array length]type_of_elems")
	arr := [3]int{1, 2, 3}
	multiDimen := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(multiDimen)

	fmt.Println("2. arrays are mutable")
	arr[1] = 4
	multiDimen[0][1] = 5
	fmt.Println(multiDimen)

	fmt.Println("3. empty arrays are init with default values:")
	var empty [2][2]int
	fmt.Println(empty)

	fmt.Println("4. ... can be used instead of size. size is set with init elements:")
	dots := [...]int{1, 4, 9}
	fmt.Println(dots)

	fmt.Println("5. arrays are VALUE, not reference types. the following does not change 'empty':")
	fmt.Println("5.5. however they CAN be passed as reference to functions to retain mutability,")
	fmt.Println("but they are passed as values by default.")
	multiDimen = empty // sizes of arrays need to match: [2][2] & [2][2]
	multiDimen[1][1] = 10
	fmt.Println(multiDimen, empty)

	fmt.Println("6. if array elems are comparable, so are the arrays:")
	fmt.Println(empty == multiDimen) // sizes need to match though
}

func copyingArray() {
	fmt.Println("Go does not have array.copy()")

	fmt.Println("option 1: manually copy from a to b:")
	var source = [2]int{1, 2}
	var destination [2]int
	for i := range len(source) {
		destination[i] = source[i]
	}
	fmt.Println(destination)

	fmt.Println("option 2: make b a POINTER to a instead:")
	var pointer *[2]int = &source
	fmt.Println(*pointer)
}
