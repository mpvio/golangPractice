package functionexamples

import "fmt"

func MethodLesson() {
	fmt.Println("methods have a 'receiver argument'")
	fmt.Println("can be a struct or not, but MUST be in same package")
	fmt.Println("cannot write method for built-in type like int")

	fmt.Println("calling a method with struct receiver:")
	structMethod()

	fmt.Println("calling a method with non-struct receiver:")
	nonStructMethod()

	fmt.Println("calling a method with pointer (i.e. reference) receiver.")
	fmt.Println("this lets variable be changed")
	fmt.Println("no need to use &, just obj.func() like usual:")
	pointerAsReceiver()

	fmt.Println("a method with value receiver (1st 2 examples)")
	fmt.Println("can be called with a point by using (&name).func")
	fmt.Println("instead of name.func as normal:")
	callValueMethodWithPointer()

	fmt.Println("all value/ reference methods can be called either way")
	fmt.Println("golang auto-converts your code to match:")
	valueAndPointerAreInterchangeable()
}

func structMethod() {
	p := person{name: "a", age: 10}
	p.display()
}

func pointerAsReceiver() {
	p := person{name: "a", age: 10}
	p.display()
	p.changeName("b")
	p.display()
}

func callValueMethodWithPointer() {
	p := person{name: "a", age: 10}
	(&p).display()
}

func valueAndPointerAreInterchangeable() {
	// fields not given are initialized with default values,
	// e.g. age = 0 here, strings would be "", etc.
	p := person{name: "abc"}

	p.display()

	p.changeName("def")
	(&p).display()

	(&p).changeName("ghi")
	p.display()
}

type person struct {
	name string
	age  int
}

func (p person) display() {
	fmt.Println("Name/ Age:", p.name, p.age)
}

func (p *person) changeName(name string) {
	p.name = name
}

func nonStructMethod() {
	n := number(5)
	o := n.square()

	fmt.Printf(`made a number 'container to write method for int. type of ans is
	%T instead of %T, but can be converted after`, o, int(o))
}

type number int

func (n number) square() number {
	return n * n
}
