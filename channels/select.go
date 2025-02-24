package channels

import (
	"fmt"
	"time"
)

func SelectBasic() {
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}

	select {
	case <-c1:
		//
	case <-c2:
		//
	case c3 <- struct{}{}:
		fmt.Println("haha")
	default:
		fmt.Println("test")
	}
}

func SelectWithChannelClosed() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		select {
		// c1 and c2 have equal chance of being selected. Go runtime will
		// perform pseudo-random uniform selection over the set of case statements
		case <-c1:
			fmt.Println("Signal received from c1")
			fmt.Println(<-c1)
			c1Count++
		case <-c2:
			fmt.Println("Signal received from c2")
			fmt.Println(<-c2)
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}

func SelectWithGoRoutine() {
	done := make(chan interface{})
	lanjut := make(chan interface{})
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		close(lanjut)
	}()

	workCounter := 0

ardhan:
	for {
		select {
		case <-done: // receive signal channel closed, let's break the ardhan
			fmt.Println("Done")
			break ardhan
		case <-lanjut:
			fmt.Println("Lanjut!")
			break ardhan
		default:
		}

		// Simulate work
		workCounter++
		fmt.Println("Nunggu signal")
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
