package main

import "flag"
import "fmt"

func main() {
    workers := flag.Uint64("workers", 0, "max integer to try")
    rng := flag.Uint64("range", 0, "max integer to try")
    flag.Parse()

    partialResults := make(chan float64, 1000000)
    quit := make(chan bool)

    go accumulateResult(*workers, partialResults, quit)

    for i := uint64(0); i < *workers; i++ {
        makeWorker((0 == ((i * *rng) % 2)), 
            float64(i * *rng), 
            float64((i+1) * *rng - 1),
            partialResults)
    }
    <- quit
} 


func makeWorker(shouldAdd bool, begin, end float64, write chan float64) {
    go func() {
        var sum float64 = 0.0
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

func accumulateResult(workers uint64, partials <- chan float64, quit chan bool) {
    var pi float64 = 0.0
    for i := uint64(0); i < workers; i++ {
        pi += <- partials
    }
    fmt.Printf("Approximate Pi = %v\n", pi)
    quit <- true
}

