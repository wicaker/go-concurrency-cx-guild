package channels

import (
	"fmt"
	"time"
)

// func BufferBasic() {
// 	var stdoutBuff bytes.Buffer
// 	defer stdoutBuff.WriteTo(os.Stdout)

// 	intStream := make(chan int, 5)
// 	go func() {
// 		defer close(intStream)
// 		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
// 		for i := 0; i < 5; i++ {
// 			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
// 			intStream <- i
// 		}
// 	}()

// 	for integer := range intStream {
// 		fmt.Fprintf(&stdoutBuff, "Received %v. \n", integer)
// 	}
// }

func BufferBasic() {
	// Create a buffered channel with a capacity of 3
	ch := make(chan int)

	// Start a goroutine to send data to the channel
	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Printf("Sending %d to the channel\n", i)
			ch <- i
		}
		// Close the channel when done sending
		close(ch)
	}()

	// Allow some time for the goroutine to start
	time.Sleep(time.Second)

	// Receive data from the channel
	for num := range ch {
		fmt.Printf("Received %d from the channel\n", num)
	}

	// Channel is closed, and all values have been received
	fmt.Println("Done receiving")
}

func BufferWithBestPractice() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)

		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()

		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
