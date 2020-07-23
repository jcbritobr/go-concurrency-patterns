package main

import "fmt"

// generator will receive a chunck of int data and return an int channel
func generator(buffer ...int) <-chan int {
	// a channel that will be returned with data
	c := make(chan int)
	// parallel process
	go func() {
		// iterate over incomming data chunk
		for item := range buffer {
			// pass data for channel
			c <- item
		}
		// when finnished, close the channel
		close(c)
	}()
	//return the channel
	return c
}

func main() {
	// data chunk
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// load data into channel
	g := generator(data...)

	// consum the channel
	for i := range g {
		// print result data chunk
		fmt.Println("item:", i)
	}

}
