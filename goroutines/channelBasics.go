package goroutines

import (
	"fmt"
	"time"
)

func ChannelInfo() {
	creatingChannels()
	sendAndReceive()
	buffering()
	closeChannelViaLoop()
	unidirectionalChannel()
}

func creatingChannels() {
	// channels are used by goroutines to communicate with each other
	// bidirectional by default (i.e a channel can send AND receive info)
	// can be declared two ways
	var chan1 chan int
	fmt.Println("this channel has no value:", chan1)

	chan2 := make(chan int)
	fmt.Println("this channel's value is its memory location:", chan2)
}

func sendAndReceive() {
	fmt.Println("channels can send/ receive data OF THEIR DECLARED TYPE (e.g. chan *int*)")
	fmt.Println("this is done with <-, e.g. 'chan <- var' or 'var <- chan'")
	fmt.Println("if there's nothing to return, just '<- chan' is also valid")
	fmt.Println("chans can be used to coordinate goroutines without needing to lock them")
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	res := <-ch
	fmt.Println(res)
	fmt.Println("close a channel with close(chan)")
	close(ch)
}

func buffering() {
	// because this has no buffering, it needs a receiver set up
	// before data is passed to it, i.e. a goroutine function
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	fmt.Println("size/ capacity of unbuffered channel before emptying:", len(ch), cap(ch))
	val := <-ch
	fmt.Println("size/ capacity of unbuffered channel aft emptying:", len(ch), cap(ch))
	close(ch)
	fmt.Println(val)

	// a N buffered channel can hold N values before deadlocking
	// (i.e. waiting for a receiver)
	ch2 := make(chan string, 3)
	ch2 <- "first"
	ch2 <- "second"
	ch2 <- "third"
	fmt.Println("size/ capacity of buffered channel before emptying:", len(ch2), cap(ch2))

	x, y, z := <-ch2, <-ch2, <-ch2
	time.Sleep(time.Second * 2)
	fmt.Println("channel values are output FIFO: x, y, z:", x, y, z)
	fmt.Println("size/ capacity of buffered channel after emptying:", len(ch2), cap(ch2))
	close(ch2)
}

func test(ch chan string) {
	for i := range 4 {
		ch <- fmt.Sprintf("HI %d\n", i)
	}
	close(ch)
}

func closeChannelViaLoop() {
	c := make(chan string)
	go test(c)

	fmt.Println("channel itself can also be iterated over:")
	for res := range c {
		fmt.Println(res)
	}

	for {
		// two part assignment: second part is bool denoting if channel is open
		res, ok := <-c
		if !ok {
			fmt.Println("ok val is true if chan is open, false if closed")
			fmt.Println("because loop runs forever, it will catch end of channel")
			break
		}
		fmt.Println("chan is open:", res, ok)
	}

}

func unidirectionalChannel() {
	// only for receiving
	chR := make(<-chan string)
	//only for sending
	chS := make(chan<- string)
	//channel types
	fmt.Printf("R: %T, S: %T\n", chR, chS)
	//receive-only channels can't be closed
	//but sending-only ones can
	close(chS)

	//converting bidirectional channel to unidirectional
	chB := make(chan string)
	go sendingOnly(chB)
	//can print directly, no need to assign to var first
	fmt.Println(<-chB)
	close(chB)

}

func sendingOnly(s chan<- string) {
	s <- `sending a bidirectional channel to a func 
	asking for unidirectional converts it to that direction.
	however, a unidirectional CANNOT be converted to bidirectional`
}
