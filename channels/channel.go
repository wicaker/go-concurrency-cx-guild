package channels

import (
	"fmt"
	"sync"
)

func BasicBidirectional() {
	var dataStream chan interface{}
	dataStream = make(chan interface{})
	dataInt := make(chan int16)

	go func() {
		dataStream <- "Hello channels"
		dataInt <- 1
	}()

	fmt.Println(<-dataStream)
	fmt.Println(<-dataInt)
}

// func BasicUnindirectionalError() {
// 	writeStream := make(chan<- interface{})
// 	readStream := make(<-chan interface{})

// 	<-writeStream
// 	readStream <- struct{}{}

// 	// unindirectional is not often we see
// 	// used as function parameters and return types mostly
// }

// func BasicUnindirectionalUsecase() {

// }

// func unindirectionalUsecase(writeStream chan<- interface{}, readStream <-chan interface{}) (chan<- interface{}, <-chan interface{}) {

// }

// func BasicUnindirectionalReceiveOrReadOnlyCorrect() {
// 	var dataStream <-chan interface{}
// 	dataStream = make(<-chan interface{})
// }

// func BasicUnindirectionalReceiveOrReadOnlyInCorrect() {
// 	var dataStream <-chan interface{}
// 	dataStream = make(<-chan interface{})
// }

// func BasicUnindirectional() {
// 	var receiveChan <-chan interface{}
// 	var sendChan chan<- interface{}
// 	dataStream := make(chan interface{})

// 	// Valid statements:
// 	receiveChan = dataStream
// 	sendChan = dataStream
// }

// func BasicUnindirectionalSendOnly() {
// 	var dataStream chan<- interface{}
// 	dataStream = make(chan<- interface{})
// }

//
func Deadlock() {
	stringStream := make(chan string)
	go func() {
		// if 0 != 1 {
		// 	return
		// }

		// stringStream <- "Hello channels!"
	}()

	fmt.Println(<-stringStream)
}

func CheckClosedChannel() {
	stringStream := make(chan string)

	go func() {
		stringStream <- "Hello channels!"
	}()

	salutation, ok := <-stringStream
	fmt.Printf("(%v): (%v)\n", ok, salutation)
	fmt.Printf("(%v): (%v)\n", ok, salutation)

	close(stringStream)
	salutation, ok = <-stringStream
	fmt.Printf("(%v): (%v)\n", ok, salutation)
}

func BestPattern() {
	numStream := make(chan int)
	go func() {
		defer close(numStream)

		for i := 1; i <= 5; i++ {
			// fmt.Println("ahaha")
			numStream <- i
		}
	}()

	for num := range numStream {
		fmt.Printf("%v ", num)
	}
}

// ?????
// Need deeper undestanding
func UnblockMultipleGoroutines() {
	begin := make(chan interface{})

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

func ClosedChannel() {
	stringStream := make(chan string, 2)

	// Send values to the channel
	go func() {
		defer close(stringStream) // Close the channel

		stringStream <- "Hello"
		stringStream <- "World"
	}()

	// for {
	ak, ok := <-stringStream
	// if !ok {
	// 	fmt.Println("Channel is closed.")
	// 	break
	// }
	fmt.Println(ak, ok)
	// }

}

func Channel() {
	stringStream := make(chan string)
	go func() {
	}()

	fmt.Println(<-stringStream)
}
