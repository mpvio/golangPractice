package testinglessons

import (
	"fmt"
	"testing"
)

func TestTestFunction(t *testing.T) {
	fmt.Println("tests are written in _test.go files")
	fmt.Println("test functions take t *testing.T as parameter")
	fmt.Println("and are only run when calling 'go test'")
	input := "s"
	expected := "s s"
	actual := TestFunction(input)
	if expected != actual {
		t.Errorf("expected (%s) != actual (%s)", expected, actual)
	}
}
