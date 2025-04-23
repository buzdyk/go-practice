package main

import "fmt"

func add(doneCh chan struct{}, inputCh chan int) chan int {
	addRes := make(chan int)

	go func() {
		defer close(addRes)

		for data := range inputCh {
			result := data + 1

			select {
			case <-doneCh:
				return
			case addRes <- result:
			}
		}
	}()
	return addRes
}

func multiply(doneCh chan struct{}, inputCh chan int) chan int {
	multiplyRes := make(chan int)

	go func() {
		defer close(multiplyRes)

		for data := range inputCh {
			result := data * 2

			select {
			case <-doneCh:
				return
			case multiplyRes <- result:
			}
		}
	}()

	return multiplyRes
}

func generator(doneCh chan struct{}, input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			select {
			case <-doneCh:
				return
			case inputCh <- data:
			}
		}
	}()

	return inputCh
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}

	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := generator(doneCh, input)

	resultCh := multiply(doneCh, add(doneCh, inputCh))

	for res := range resultCh {
		fmt.Println(res)
	}
}
