package main

import "go-concurrency/mutexracefix"

// import "go-concurrency/channels"

// import "go-concurrency/goroutines"

func main() {
	// goroutines.Goroutines()
	// goroutines.GoroutinesAnonymous()
	// goroutines.GoroutinesAnonymousWithVariableDeclared()
	// goroutines.GoroutinesJoinPoint()
	// goroutines.Closures()
	// goroutines.ClosuresWithArray()
	// goroutines.ClosuresWithArray2()
	// goroutines.MemoryAllocation()
	// goroutines.MutexBasic()
	// goroutines.RaceCondition() // go run -race main.go
	// channels.BasicBidirectional()
	// channels.BasicUnindirectionaError()
	// channels.Deadlock()
	// channels.CheckClosedChannel()
	// channels.BestPattern()
	// channels.UnblockMultipleGoroutines()
	// channels.ClosedChannel()
	// channels.BufferBasic()
	// channels.BufferWithBestPractice()
	// channels.SelectBasic()
	// channels.SelectWithChannelClosed()
	// channels.SelectWithGoRoutine()
	mutexracefix.RaceConditionFix() // go run -race main.go

	// stringStream := make(chan string)
	// go func() {
	// 	// stringStream <- "jawaban"
	// }()

	// fmt.Println(<-stringStream)

	// go server()

	// var wg sync.WaitGroup

	// // dt := time.Now()

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go server(&wg, i)
	// 	// wg.Done()
	// }

	// wg.Wait()
}

// func server(wg *sync.WaitGroup, i interface{}) {
// 	defer wg.Done()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(fmt.Sprintf("server %d", i))
// }
