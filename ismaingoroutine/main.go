package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Goroutine num includes main processing
	fmt.Println(runtime.NumGoroutine()) // 1

	// Spawn two goroutines
	go func() { time.Sleep(1 * time.Second) }()
	go func() { time.Sleep(1 * time.Second) }()

	// Total three goroutine runs
	fmt.Println(runtime.NumGoroutine()) // 3
}
