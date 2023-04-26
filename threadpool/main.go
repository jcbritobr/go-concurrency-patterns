package main

import (
	"sync"
)

type ExecFunc func()

type ThreadPool struct {
	receiver chan ExecFunc
	size     int
}

func NewThreadPool(size int) *ThreadPool {
	pool := &ThreadPool{
		receiver: make(chan ExecFunc),
		size:     size,
	}

	return pool
}

func (t ThreadPool) start() {
	for i := 0; i < t.size; i++ {
		go func() {
			for {
				f := <-t.receiver
				f()
			}
		}()
	}
}

func (t *ThreadPool) Execute(f ExecFunc) {
	t.receiver <- f
}

func main() {
	pool := NewThreadPool(3)
	var wg sync.WaitGroup
	wg.Add(2)
	pool.start()

	job := func(id int) {
		for i := 0; i < 1000; i++ {
			println(id, "->", i)
		}
	}

	pool.Execute(func() {
		job(1)
		wg.Done()
	})

	pool.Execute(func() {
		job(2)
		wg.Done()
	})

	wg.Wait()
}
