package basic

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func MyErrorHandler() {
	num := -1
	res, err := throwError(num)
	if err != nil {
		// error handling
		fmt.Println("An error was thrown:", err)
		// can panic instead
		// panic(err)
		return
	}
	fmt.Println(res)
}

func FileIOExample() {
	createFile()
	readFile()
}

func TakeUserInput() {
	var input string
	fmt.Println("user text is written to 'input' via reference (and can accept int values as well):")
	/* note: scanln doesn't work when running code from vs code
	(needs to be run in terminal) */
	fmt.Scanln(&input)
	fmt.Printf("you typed: %s\n", input)
}

func WithBufioScanner() {
	fmt.Println("for simpler inputs")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()         // start reading
	line := scanner.Text() // get input
	fmt.Println(line, "no need for delimiter")
}

func WithBufioReader() {
	fmt.Println("for more complex inputs, e.g. runes, peeking ahead")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	fmt.Println(line, "need to specify delimiter (end of line character), which is included at the end of the 'line'")
}

func GetWorkingDirectory() {
	mydir, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(mydir)
}

// error is passed as final return param per tradition
func throwError(x int) (int, error) {
	if x < 0 {
		// 0 here is a default value
		return 0, errors.New("write an error msg here")
	}
	// nil here is an absent error
	return x * x, nil
}

func createFile() {
	fmt.Println("writing to test.txt")
	file, err := os.Create("test.txt")

	if err != nil {
		panic(err)
	}

	length, err := file.WriteString("write strings" + ", these are all one line" + "use file.Write() for bytes")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File name: %s, length of text: %d\n", file.Name(), length)
}

func readFile() {
	fmt.Println("reading test.txt")
	fileName := "test.txt"
	data, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Printf("file size: %d\n", len(data))
	fmt.Printf("file content: %s\n", data)
}
