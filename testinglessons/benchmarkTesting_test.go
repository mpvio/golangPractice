package testinglessons

import (
	"fmt"
	"testing"
)

func BenchmarkTestTestFunction(b *testing.B) {
	fmt.Println("used to test performance/ spd of func,")
	fmt.Println("in/ correct results don't matter:")
	fmt.Println("run with go test -bench= (can follow with extra flags if needed)")
	TestFunction("hello")
}
