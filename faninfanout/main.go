package main

import (
	"fmt"
	"sync"
)

// generator transforms and transports data from
// buffer to other channels
func generator(buffer ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, item := range buffer {
			c <- item
		}
		close(c)
	}()

	return c
}

// sq implements a square pipeline
func sq(buffer <-chan int) <-chan int {

	c := make(chan int)
	go func() {
		for i := range buffer {
			c <- i * i
		}
		close(c)
	}()

	return c
}

// merge merge all channels in one - fan in
func merge(buffer ...<-chan int) <-chan int {
	// create sync mechanism and an out channel
	var wg sync.WaitGroup
	out := make(chan int)

	// create a function to merge the channels
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	// add the size of workers to waitgroup
	wg.Add(len(buffer))
	// merge the channels
	for _, c := range buffer {
		go output(c)
	}

	// this function will wait the other goroutines and close the out channel
	go func() {
		wg.Wait()
		close(out)
	}()

	// return the results
	return out
}

func main() {
	// data to be processed
	data := []int{1, 2, 3, 4, 5}
	// generator loads data
	g := generator(data...)

	// fan out the data in square pipeline
	sq1 := sq(g)
	sq2 := sq(g)

	// consumn the parallel square pipeline
	for item := range merge(sq1, sq2) {
		fmt.Println(item)
	}
}
