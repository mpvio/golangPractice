package goroutines

import (
	"fmt"
	"sync"
)

var num = 0

func sampleFunction(w *sync.WaitGroup, m *sync.Mutex) {
	// this part of code is locked
	m.Lock()
	// anything between lock and unlock
	// can only be executed by one goroutine at a time
	num += 1
	m.Unlock()
	// if this func is all the routine has to do,
	// we can decrement the wait counter by 1
	// waiting ends once counter is 0
	w.Done()
}

func Introduction() {
	// mutex locks are done to ensure only 1 goroutine accesses a locked
	// length of code
	var m sync.Mutex
	// when a waitgroup waits, it only continues once
	// all goroutines in it are complete (so don't need to use time.sleep())
	var w sync.WaitGroup

	for range 10 {
		//adding one goroutine so add 1 to counter
		w.Add(1)
		go sampleFunction(&w, &m)
	}

	// now code will wait until counter goes back down to 0
	w.Wait()
	fmt.Println("Result:", num)
}
