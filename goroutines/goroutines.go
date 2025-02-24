package goroutines

import (
	"fmt"
	"runtime"
	"sync"
)

// goroutine & channel are golang concurrency primitives

// Groutine
// Higher level of abstraction known as coroutines

func Goroutines() {
	go sayHello()

	fmt.Println("World")
}

func sayHello() {
	fmt.Println("Hello!")
}

func GoroutinesAnonymous() {
	go func() {
		fmt.Println("Hello")
	}()

	fmt.Println("World")
}

func GoroutinesAnonymousWithVariableDeclared() {
	sayHello := func() {
		fmt.Println("Hello")
	}

	go sayHello()

	fmt.Println("World")
}

// simple examples
// now we can print inside the go routine function
func GoroutinesJoinPoint() {
	var wg sync.WaitGroup // digunakan untuk menunggu goroutine

	sayHello := func() {
		// Done() is called within each goroutine to signal that it has completed its work.
		// When a goroutine finishes its task, it should call Done() to decrement the internal counter in the WaitGroup.
		// When the counter reaches zero, any goroutine waiting on Wait() will unblock.
		defer wg.Done() // defer keyword is used to delay the execution of a function or a statement until the nearby function returns

		fmt.Println("Run go routines with GoroutinesJoinPoint")
	} // this line seem like never being executed

	// You call Add(n) to indicate that you expect to wait for n goroutines to finish their work.
	wg.Add(1) // is used to specify the number of goroutines you want to wait for.

	go sayHello()

	// blocks the execution of the current goroutine until the internal counter in the WaitGroup reaches zero.
	// It effectively waits for all the goroutines you've indicated with Add() to call Done() and signal that they have completed their work.
	// Once the counter reaches zero, Wait() unblocks, and your program can continue.
	wg.Wait()

	fmt.Println("Outside the function")
}

func Closures() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)

	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()

	wg.Wait()

	fmt.Println(salutation)
}

func ClosuresWithArray() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}

	wg.Wait()
}

func ClosuresWithArray2() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}

	wg.Wait()
}

func MemoryAllocation() {
	// only 2-4 kb for empty goroutines
	// go func() {
	// }()

	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}

	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }

	const numGoroutines = 10000
	wg.Add(numGoroutines)
	before := memConsumed()

	for i := numGoroutines; i > 0; i-- {
		go noop()
	}

	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}

// now we can print inside the go routine function
func GoroutinesWithWaitGroup() {
	var wg sync.WaitGroup // digunakan untuk menunggu goroutine

	sayHello := func() {
		// Done() is called within each goroutine to signal that it has completed its work.
		// When a goroutine finishes its task, it should call Done() to decrement the internal counter in the WaitGroup.
		// When the counter reaches zero, any goroutine waiting on Wait() will unblock.
		defer wg.Done() // defer keyword is used to delay the execution of a function or a statement until the nearby function returns

		fmt.Println("Run go routines with GoroutinesWithWaitGroup")
	}

	// You call Add(n) to indicate that you expect to wait for n goroutines to finish their work.
	wg.Add(1) // is used to specify the number of goroutines you want to wait for.

	go sayHello()

	// blocks the execution of the current goroutine until the internal counter in the WaitGroup reaches zero.
	// It effectively waits for all the goroutines you've indicated with Add() to call Done() and signal that they have completed their work.
	// Once the counter reaches zero, Wait() unblocks, and your program can continue.
	wg.Wait()

	fmt.Println("Outside the function")
}
