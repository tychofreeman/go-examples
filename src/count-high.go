package main

import "fmt"

func worker(id int, nums <-chan int, quit chan int) {
	fmt.Printf("Starting worker %v...\n", id)
	for num := range nums {
		fmt.Printf("%v: %v\n", id, num)
	}
	quit <- id
	fmt.Printf("Workder %v done\n", id)
}

const MAX = 100000

func main() {
	quit := make(chan int, 20)
	for i := 0; i < MAX; i++ {
		nums := make(chan int, 30)
		go worker(i, nums, quit)
		for j := 0; j < 1; j++ {
			nums <- j
		}
		close(nums)
	}
	fmt.Printf("Done setting up workers...\n")

	for i := 0; i < MAX; i++ {
		id := <- quit
		fmt.Printf("Worker %v quit\n", id)
	}
}
