package basic

import "fmt"

func ConditionalsShowcase() {
	fmt.Println("If Examples:")
	ifExamples(100)
	ifExamples(200)
	ifExamples(300)

	fmt.Println("For Examples:")
	forLoopExamples()

	fmt.Println("For loop over channel:")
	forLoopOverChannel()

	fmt.Println("showcase continue and GoTo (break used to break infinite loop earlier)")
	goToAndContinue()
}

func ifExamples(x int) {

	if x < 200 {
		if x == 100 {
			fmt.Println("This is a nested if statement.")
		}
	} else if x == 200 {
		fmt.Println("this is an else-if/ elif check")
	} else {
		fmt.Println("this is an else")
	}
}

func forLoopExamples() {

	fmt.Println(`for loops take the following form:
	'for intialization; condition; [what to do after each iteration] {}'`)

	for i := 0; i < 4; i += 2 {
		fmt.Printf("i is %d\n", i)
	}

	// this is shorthand for: for i := 0; i < 4; i ++ (i.e. 0, 1, 2, 3 <- step = 1)
	for j := range 4 {
		fmt.Printf("j is %d\n", j)
	}

	for {
		fmt.Println("with no conditions, this loop is infinite. escape it with break")
		break
	}

	j := 0
	for j < 3 {
		fmt.Println("Go has no while loop, but using a boolean for the condition mimics one")
		j += 2
	}

	ls := []string{"this", "is", "a", "list"}
	fmt.Println(`define a foreach loop as 'for index, value := range VARIABLE_NAME
	for x := range will default to x as index
	can use _ to ignore unwanted side, e.g. _, value to just get values`)
	for i, val := range ls {
		fmt.Println(i, val)
	}

	fmt.Println("can iterate over strings and maps as well")

	str := "test"
	for i, j := range str {
		//%U is unicode value
		fmt.Printf("The index number of %U is %d\n", j, i)
	}

	mmap := map[int]string{
		1: "[key type]",
		2: "next is value type",
		3: "need a comma after every line, even the last!",
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}
}

func forLoopOverChannel() {
	//create channel
	chn1 := make(chan int)

	go func() {
		//assign these values to channel (like adding to a list)
		chn1 <- 100
		chn1 <- 1000
		//close channel once finished
		close(chn1)
	}() //() at the end as if the {} is its own function

	for i := range chn1 {
		fmt.Println("i is int type, value:", i)
	}
}

func goToAndContinue() {

	for i := range 3 {
		if i == 2 {
			fmt.Println("Continue is used to skip a loop iteration, e.g. if i == 2")
			continue
		}
		fmt.Println(i)
	}

	x := 0

label1:
	for x < 8 {
		if x == 5 {
			fmt.Println("use goto to jump to a defined label, here go back to the top of the loop when x == 5")
			x += 2
			goto label1
		}
		fmt.Print(x, " ")
		x++
	}
	fmt.Println("")
}

func Switches() {
	fmt.Println("Two types: expression and type switches")
	fmt.Println("Expression switches can have optional statements and expressions")

	fmt.Print("Optional statement: ")
	// statement; the expression the switch checks (here it's x)
	switch x := 3; x {
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}

	fmt.Print("Optional expression: ")
	// statement; what switch checks
	x := 3
	switch {
	case x == 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}

	fmt.Print("can combine both: ")
	// statement; what switch checks
	switch x := 3; {
	case x == 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}

	fmt.Print("type switch used to branch on interface type, not value: ")
	fmt.Println("e.g. for generic functions")
	var number interface{} = 4
	// var number any = 4 <- this is also valid
	switch t := number.(type) {
	case int:
		fmt.Println("number")
	default:
		fmt.Println("other", t)
	}
}

func SelectStatement() {
	fmt.Println("select is similar to switch, but cases refer to send/ receive operations on channels")

	c := make(chan int)

	select {
	case <-c:
	default:
		fmt.Println("'case <- c: {None}' does nothing, so having a default to fall back on avoids deadlock")
	}

	select {
	case x1 := <-c:
		fmt.Println("value: ", x1)
	default:
		fmt.Println("the above doesn't send/ receive anything so is 'nil'. thus default is called even though the above case has code in it")
	}
}
