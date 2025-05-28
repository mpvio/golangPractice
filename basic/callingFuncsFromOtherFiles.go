package basic

import "fmt"

func Multiply(a, b int) int {
	return a * b
}

func HowToUseMultipleGoFiles() {
	launchJsonSettingForMultipleGoFiles()
	otherModule()
}

func launchJsonSettingForMultipleGoFiles() {
	fmt.Println("")
	fmt.Println("If using more than one file, set 'program' to '.' in launch.json")
	fmt.Println("this translates to 'go run .' which is needed to include all files, not just 'main.go'")
}

func otherModule() {
	fmt.Println("")
	fmt.Println("IT SHOULDN'T BE THIS ANNOYING TO SET UP IMPORTS!")
	fmt.Println("create go.mod file with module name (e.g. self) and go version")
	fmt.Println("each folder needs its own package name")
	fmt.Println("(name is irrelevant, but it should be the same across all files in that folder)")
	fmt.Println("when calling package, use module name/file path starting from where go.mod is")
	fmt.Println("e.g. here, the /basic folder is in the folder go.mod is, so: 'self/basic'")
	fmt.Println("can prefix it with an alias to make calling functions from there easier, e.g. 'b self/basic'")
	fmt.Println("so I can call 'b.Multiply' below instead of 'basic.Multiply', e.g:")
	fmt.Println(Multiply(1, 2))
}
