package functionexamples

import "fmt"

type student struct {
	name  string
	grade int
}

type teacher struct {
	name   string
	degree float32
}

func (x student) show() {
	fmt.Println("Name:", x.name)
}

func (x teacher) show() {
	fmt.Println("Name:", x.name)
}

func StructMethodOverloading() {
	s := student{name: "a", grade: 10}
	t := teacher{name: "b", degree: 10.0}

	fmt.Println("two structs can have funcs with the same name:")
	s.show()
	t.show()
}

type v1 string
type v2 int

func (x v1) show() v1 {
	return "a" + x
}

func (x v2) show() v2 {
	return 5 + x
}

func NonStructMethodOverloading() {
	a := v1("Hello")
	b := v2(10)

	fmt.Println("two non-structs can have funcs with the same name:")
	fmt.Println(a.show())
	fmt.Println(b.show())
}
