package main

import (
	"fmt"
	"flag"
)

func printPrimes(primes chan chan int) {
	for result := range primes {
		for p := range result {
			if p != 0 {
				fmt.Printf("%v, ", p)
			}
		}
	}
}

func isPrime(i int, primes chan chan int) {
	myPrime := make(chan int)
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

func main() {
	max := flag.Int("max", 0, "max integer to try")
	flag.Parse()
	fmt.Printf("Got arg %v\n", *max)

	primes := make(chan chan int, 10)

	go printPrimes(primes)

	
	isPrime(2, primes)
	for i := 3; i < *max; i = i + 2 {
		isPrime(i, primes)
	}
	close(primes)
}
