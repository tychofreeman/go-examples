package main

import (
	"fmt"
)

func generate() (ch chan uint64) {
	ch = make(chan uint64)
	go func() {
		for i := uint64(2); ; i++ {
			ch <- i
		}
	}()
	return
}

func filter(in chan uint64, prime uint64) (out chan uint64) {
	out = make(chan uint64)
	go func() {
		for {
			if i := <- in; i % prime != 0 {
				out <- i
			}
		}
	}()
	return
}

func sieve() (out chan uint64) {
	out = make(chan uint64)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			out <- prime
			ch = filter(ch, prime)
		}
	}()
	return
}

func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}
