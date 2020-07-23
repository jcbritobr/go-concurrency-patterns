## Golang Concurrent Patterns Examples
This repository contains golang concurrent patterns implemented examples.

- Fan in / Fan out \
The best way to implement a processing pipeline in golang is using the fan in/out pattern. The pattern is built by a function and a goroutine
that transports and load data, using channels, to another function and goroutine(pipeline). In the end o process, all the goroutines data are merged into one.
See image below: 

<div align="center">![alt](faninfanout/images/faninfanout.png)</div>

- Workers Pool