package main

import (
	"fmt"
	"flag"
)

func makeWorker(beginIsEven bool, begin, end float64, write chan float64) {
	go func() {
		var sum float64 = 0.0
		shouldAdd := beginIsEven
		for i := begin; i <= end; i = i + 1.0 {
			addend := 4.0 / (2*i + float64(1.0))
			if shouldAdd {
				sum += addend
			} else {
				sum -= addend
			}
			shouldAdd = !shouldAdd
		}
		write <- sum
	}()
}

func main() {
	workers := flag.Uint64("workers", 0, "max integer to try")
	rng := flag.Uint64("range", 0, "max integer to try")
	flag.Parse()

	chans := make(chan float64, 1000000)
	quit := make(chan bool)

	go func() {
		var pi float64 = 0.0
		for i := uint64(0); i < *workers; i++ {
			pi += <- chans
		}
		fmt.Printf("Approximate Pi = %v\n", pi)
		quit <- true
	}()

	var i uint64 = 0
	for ; i < *workers; i++ {
		makeWorker((0 == ((i * *rng) % 2)), float64(i * *rng), float64((i + 1) * *rng - 1), chans)
	}
	<- quit
}
