package main

import (
	"fmt"
	"flag"
)

func addFilter(num int, read chan int) (write chan int) {
	if closed(read) {
		return
	}

	write = make(chan int)
	
	fmt.Printf("%v\n", num)

	go func() {
		defer func() {
			close(write)
		}()

		for i := range read {
			if i < num || i % num != 0 {
				write <- i
			}
		}
	}()
	return
}

func main() {
	max := flag.Int("max", 0, "max integer to try")
	flag.Parse()
	fmt.Printf("Got arg %v\n", *max)

	in := make(chan int)

	go func() {
		for i := 2; i < *max; i++ {
			in <- i
		}
		close(in)
	}()

	out := addFilter(2, in)

	for out != nil && !closed(out) {
		prime := <- out
		out = addFilter(prime, out)

	}
}
