package main

import (
	// these need to match FOLDER NAMES
	// NOT actual package name
	// BUT package name should be the same
	// for all files in a folder
	"fmt"
	a "self/arraysAndSlices"
	b "self/basic"
	f "self/functionExamples"
	ge "self/genericsTest"
	g "self/goroutines"
	j "self/jsonHandling"
	p "self/pointers"
	r "self/reflections"
)

const PI = 3.14
const STRING = "Constants don't need to be defined WITH ':=' because they need 'const'. They can't be overwritten after."

func ErrorsAndIO() {
	b.MyErrorHandler()
	b.FileIOExample()
	//b.TakeUserInput()
	b.GetWorkingDirectory()
}

func Goroutines() {
	g.GoDisplay()
	g.ChannelInfo()
	g.Introduction()
}

func JsonHandling() {
	j.MarshalDataToJson()
	j.UnmarshalJsonToData()
}

// ever executable neds a main package and main() with no input/ output
// this is run by default when executing
func main() {
	// fmt.Println("Hello World!")
	// fmt.Println("Constants:", PI, STRING)
	JsonHandling()

}

func DataTypes() {
	fmt.Println(`data types:
		basic: numbers, strings, booleans
		aggregate: array, structs
		reference: pointers, slices, maps, functions, channels
		interface: self-explanatory`,
	)

	TypeAssertions()
	TypeSwitches()

}

func OverloadingExamples() {
	f.StructMethodOverloading()
	f.NonStructMethodOverloading()
}

func PastLessons() {
	// fmt.Println(`Can Call 'HowTo...' because starts with capital letter.
	// function names starting with lower case are private so can't be called from
	// another file`)
	b.HowToUseMultipleGoFiles()
	b.StringExamples()
	b.ConditionalsShowcase()
	b.Switches()
	b.SelectStatement()
	f.TypesOfCalls()
	f.VariadicFunction()
	f.MethodLesson()
	f.DeferenceShowcase()
	f.MethodLesson()
	OverloadingExamples()
	a.SliceInfo()
	p.PointerLesson()
	ErrorsAndIO()
	DataTypes()
	Goroutines()
	ge.UseGenericFunction()
	r.ReflectionsOne()
}

func HandleMultipleReturns() {
	var x, y = f.MultipleReturns()
	// _ means second value is ignored below:
	var a, _, c = f.ReturnNamesVals()
	fmt.Println(x, y, a, c)
}

func init() {
	fmt.Println("there can be many init functions")
}

func init() {
	fmt.Println("there can be multiple inits across the project")
}

func init() {
	fmt.Println(`they are run in alphabetical order of the files, 
	then in order of definition within them (main's are run last).
	inits are run before main() even if defined after it in that file`)
}

func init() {
	fmt.Println("this runs before main even though it was declared after.")
}
