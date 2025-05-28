package interfacework

import (
	"fmt"
	"math"
)

// 1 interface implementation
func ShapeAreas() {
	c := Circle{radius: 5}
	s := Square{side: 5}

	getArea(c)
	typeSwitching(s)

	// 6 polymorphism
	// because both shapes implement AreaInterface, they can be
	// treated as such even though they differ otherwise
	// (and implement the interface differently)
	// e.g. put in a data structure together:
	sli := []AreaInterface{c, s}
	fmt.Println(sli)
}

// 2 dynamic values and types for interfaces
func InterfaceValueAndTypes() {
	var s AreaInterface
	fmt.Println("can't create a Shape struct, but can define one")
	fmt.Printf("its value and type are null though: %s, %T\n", s, s)
}

// 3 type assertion
func getArea(s any) {
	fmt.Println(`can ask for an input of anything implementing any interface
	with 's interface{}' as param, but can also just use 's any'
	because 'any' is ALSO an interface in Go`)
	value, ok := s.(AreaInterface)
	if ok {
		fmt.Println(`
		input is identified as using the interface Shape, 
		so can use its functions AS THE RETURNED 'value', not the initial one!:
		`, value.Area())
	} else {
		fmt.Println("if ok is false, input doesn't implement Shape.")
	}
	fmt.Println(`type extraction takes the form:
	value (, ok is optional) := interfaceVariable.(InterfaceName)`)
}

// 4 type switching
func typeSwitching(s any) {
	fmt.Println("use type switch to get base type of interface variable:")
	switch obj := s.(type) {
	case Circle:
		fmt.Println("circle found:", obj.Area())
	case Square:
		fmt.Println("square found:", obj.Area())
	default:
		fmt.Println("unknown shape")
	}
}

// 5 multiple interfaces
func MultiInterfacing() {
	sqr := Square{side: 10}

	// var can be cast to any interface to use funcs specific to it when needed:
	// only if var is valid for it (e.g. has the params needed in implementation)
	var ai AreaInterface = sqr
	var pi PerimeterInterface = sqr
	fmt.Println(ai.Area(), pi.Perimeter())

	// can also declare an interface that embeds multiple to use ALL their funcs:
	// changes to areainterface and perimeterinterface will reflect in shapeinterface
	var si ShapeInterface = sqr
	fmt.Println(si.Area(), si.Perimeter())
}

type AreaInterface interface {
	//list of function signatures, no implementations
	Area() float32
}

type PerimeterInterface interface {
	Perimeter() float32
}

type ShapeInterface interface {
	AreaInterface
	PerimeterInterface
}

func (s Square) Perimeter() float32 {
	return s.side * 4
}

type Circle struct {
	radius float32
}

type Square struct {
	side float32
}

// implement an interface on a struct by defining the function FOR that struct
func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

// can implement same interface for multiple structs
func (s Square) Area() float32 {
	return s.side * s.side
}
