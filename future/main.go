package main

import "fmt"

func fiboFuture(number int) <-chan int {
	// Channel with the result
	c := make(chan int)

	// Starts a parallel process to calc fibonacci number
	go func() {
		if number <= 1 {
			c <- number
		}

		fib := 1
		prevFib := 1
		for i := 0; i < number; i++ {
			temp := fib
			fib += prevFib
			prevFib = temp
		}
		// Send the result to channel
		c <- fib
		// Close channel
		close(c)
	}()

	return c
}

func main() {
	// Get a future processing
	future := fiboFuture(44)
	// Prints the result of future. If still processing, instead, future will block
	fmt.Println("result of future processing:", <-future)

}
