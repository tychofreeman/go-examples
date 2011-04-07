package main

import (
	"fmt"
	"flag"
	. "math"
)

func makeWorker(begin, end uint64, write chan float64) {
	go func() {
		defer close(write)

		var sum float64 = 0.0
		for i := begin; i <= end; i++ {
			addend := 4.0 * Pow(-1.0, float64(i)) / float64(2 * i + 1)
			sum += addend
		}
		write <- sum
	}()
}

func main() {
	workers := flag.Uint64("workers", 0, "max integer to try")
	rng := flag.Uint64("range", 0, "max integer to try")
	flag.Parse()
	fmt.Printf("Workers: %v\nRange: %v\nTotal Iterations: %v\n", *workers, *rng, *workers * *rng)

	chans := make(chan chan float64, 100000)
	quit := make(chan bool)

	go func() {
		var pi float64 = 0.0
		for c := range chans {
			pi += <- c
			fmt.Printf("Approximate pi = %v\n", pi)
		}
		fmt.Printf("Pi = %v\n", pi)
		quit <- true
	}()

	var i uint64 = 0
	for ; i < *workers; i++ {
		write := make(chan float64)
		makeWorker(i * *rng, (i + 1) * *rng - 1, write)
		chans <- write
	}
	close(chans)
	<- quit
}
