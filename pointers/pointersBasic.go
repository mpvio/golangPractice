package pointers

import "fmt"

func PointerLesson() {
	pointers1()
	structPointers()
	ptrCapAndLen()
}

func pointers1() {

	var x *int
	fmt.Println("this is only a POINTER to an int value, currently <niL>:", x)

	y := 10
	x = &y

	fmt.Println("y's location in memory:", x)
	fmt.Println("y's value by dereferencing x:", *x)

	*x += 100
	fmt.Println("y's value after adding to '*x':", y)

	z := &x
	fmt.Printf("Type of a pointer To a pointer is **Type: %T\n", z)

	fmt.Println("pass a pointer as param to func by prefixing variable with &")
	pointerFunction(&y)
	//the below two are equivalent to the above
	//pointerFunction(x)  //is already a pointer
	//pointerFunction(*z) //dereferencing a double pointer to get initial pointer

	fmt.Println("pointers can be compared (but not ptr to ptr-ptr):", x == *z, &x != z)
}

func pointerFunction(a *int) *int {
	fmt.Println("this takes a POINTER TO an int, not an int itself:", *a)
	fmt.Println("It is also possible to return a pointer.")

	b := 50
	return &b
}

type Person struct {
	name string
}

func structPointers() {
	fmt.Println("two ways of creating a pointer to a struct")

	fmt.Println("1. use & to get address of existing struct like normal:")
	p := Person{}
	pp := &p

	fmt.Println("2. use new function. this creates a struct but returns a pointer to it:")
	ppp := new(Person)

	fmt.Println("can access struct fields directly from ptr: go auto-dereferences it:")
	pp.name = "a"
	ppp.name = "b"
	fmt.Println("first is printed as reference (&) since it isn't dereferenced:", pp, *ppp)
}

func ptrCapAndLen() {
	fmt.Println("same as for slices, but calls to ptr will check reference:")

	var arr [8]int
	arrPtr := &arr

	var nullPtr [5]*int //null ptr's cap & len will be the max size set here
	fmt.Println("arrPtr:", cap(arrPtr), len(arrPtr))
	fmt.Println("nullPtr:", cap(nullPtr), len(nullPtr))
}
