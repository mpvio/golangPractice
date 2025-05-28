package reflections

import (
	"fmt"
	"reflect"
)

type customString string
type complexStruct struct {
	name string
	age  int
}

func ReflectionsOne() {
	fmt.Println("uses reflect package,")
	fmt.Println("comes in three types: valueOf, typeOf, kind:")

	name := customString("Name")
	fmt.Println("Type is the user defined type name:", reflect.TypeOf(name))
	fmt.Println("Value is the value of the variable:", reflect.ValueOf(name))
	fmt.Println("Kind is the ACTUAL type name:", reflect.ValueOf(name).Kind())

	person := complexStruct{}
	fmt.Println("Person type:", reflect.TypeOf(person))
	fmt.Println("Person value:", reflect.ValueOf(person))
	fmt.Println("Person kinda:", reflect.ValueOf(person).Kind())

	fmt.Println("can further break structs down:")
	value := reflect.ValueOf(person)
	fields := value.NumField() // only valid if Kind == reflect.struct
	for i := range fields {
		v := value.Field(i) // of type reflect.Value, only value value == reflect.struct!
		fmt.Printf("%d. Type: %T, Value: %#v\n", i, v, v)
	}
}
