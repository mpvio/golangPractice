package goroutines

import (
	"fmt"
	"time"
)

func display() {
	str := "abcd"
	for i := range str {
		// use time.sleep to ensure code waits for this to complete
		// before terminating
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str[i])
	}
}

func GoDisplay() {
	// goroutines: less memory than threads
	go display() // runs this in parallel.
	fmt.Println("Hello World!")

	//goroutines can also be run off of anonymous functions
	go func() {
		for i := range 3 {
			fmt.Println("Anon", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second) //wait for routines to finish (not guaranteed to ALL finish)
}
