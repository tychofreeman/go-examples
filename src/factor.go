package main

import (
	"fmt"
	"flag"
)

func main() {
	max := flag.Int("max", 0, "max integer to try")
	flag.Parse()
	fmt.Printf("Got arg %v\n", *max)

	primes := make(chan chan int, 10)

	go func() {
		for result := range primes {
			for p := range result {
				if p != 0 {
					fmt.Printf("%v, ", p)
				}
			}
		}
	}()

	for i := 2; i < *max; i++ {
		myPrime := make(chan int, 2)
		primes <- myPrime
		num := i
		go func() {
			defer close(myPrime)
			for j := 2; j*j <= num ; j++ {
				if num % j == 0 {
					myPrime <- 0
					return
				}
			}
			myPrime <- num
		}()
	}
	close(primes)
}
