package main

import (
	"fmt"
	"sync"
	"time"
)

// generator 生成一系列数据并发送到通道
func generator(start int, end int, c chan<- int) {
	for i := start; i <= end; i++ {
		c <- i
		time.Sleep(time.Millisecond * 500) // 模拟延迟
	}
	close(c)
}

// fanIn 将多个通道合并为一个通道
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _, c := range channels {
		wg.Add(1)
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func Test() {
	c1 := make(chan int)
	c2 := make(chan int)

	go generator(1, 5, c1)
	go generator(6, 10, c2)

	for n := range fanIn(c1, c2) {
		fmt.Println(n)
	}
}
