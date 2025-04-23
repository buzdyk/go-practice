package main

import (
	"fmt"
)

func generator(numbers []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range numbers {
			out <- n
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()

	return out
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create pipeline
	source := generator(numbers)
	results := square(source)

	for result := range results {
		fmt.Println(result)
	}
}
