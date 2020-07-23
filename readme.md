## Golang Concurrent Patterns Examples
This repository contains golang concurrent patterns implemented examples.

* Generator/Iterator \
Generator pattern is a way to a generate data sequences using paralelism. Using this pattern we are able to make the consumer run in parallel with the generator.

**Implementation**
```go
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
```

**Use**
```go
	// data chunk
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// load data into channel
	g := generator(data...)

	// consum the channel
	for i := range g {
		// print result data chunk
		fmt.Println("item:", i)
	}
```
* Future \
A Future will start a computation in parallel, and its results will be available in the future. In golang, a simple goroutine can be use to implement this funcionality, without no use of third party or even standard libraries. 

* Fan-in/Fan-out \
The best way to implement a processing pipeline in golang is using the fan in/out pattern. The pattern is built by a function and a goroutine
that transports and loads data, using channels, to another. In the end of process, all the goroutines data are merged into one. Its a way of multiplexing and demultiplexing multiple input data.
See image below:

<p align="center">
    <img src="faninfanout/images/faninfanout.png">
</p>

* Workers Pool
