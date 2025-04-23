package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := generator("Hello")
	for msg := range ch {
		fmt.Println(msg)
	}
}

func generator(s string) chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := 0; i < len(s); i++ {
			ch <- s + " " + strconv.Itoa(i)
		}
	}()

	return ch
}
