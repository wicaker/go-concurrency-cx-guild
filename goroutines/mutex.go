package goroutines

import (
	"fmt"
	"runtime"
	"sync"
)

// A Mutex provides a concurrent-safe way to express exclusive
// access to these shared resources.
func MutexBasic() {
	var mu sync.Mutex
	sharedResource := 0

	// Simulate some work
	for i := 0; i < 5; i++ {
		// Acquire the lock to access the shared resource
		mu.Lock()

		// Modify the shared resource
		sharedResource++
		fmt.Printf("Iteration %d: sharedResource = %d\n", i+1, sharedResource)

		// Release the lock
		mu.Unlock()
	}

	fmt.Printf("All iterations have finished. Final sharedResource = %d\n", sharedResource)
}

func MutexIncDec() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}

type counter struct {
	val int
}

func (c *counter) Add(int) {
	c.val++
}

func (c *counter) Value() int {
	return c.val
}

func RaceCondition() {
	runtime.GOMAXPROCS(2) // sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting

	var wg sync.WaitGroup
	var meter counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}
