package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

type Name string

func (n Name) Hello() error {
	if n == "Gleb" {
		return errors.New("I don't know Gleb")
	}
	fmt.Println("Hello ", n)
	return nil
}

var names = []Name{"Anna", "Ivan", "Fedor", "Katya", "Gleb"}

func main() {
	g := &errgroup.Group{}
	var name Name
	for _, name = range names {
		g.Go(name.Hello)
	}

	err := g.Wait()

	fmt.Println("Error", err)
}
