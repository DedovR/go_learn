package main

import (
	"fmt"
	"sync"
	"time"
)

func mergeChannels(chs ...<-chan int) <-chan int {
	outCh := make(chan int)

	go func(chs []<-chan int, out chan int) {
		wg := &sync.WaitGroup{}
		for _, ch := range chs {
				wg.Add(1)
				go func(wg *sync.WaitGroup, c <-chan int, out chan int)  {
						for m := range c {
								out <- m
						}
						wg.Done()
				}(wg, ch, outCh)
		}
		wg.Wait()

		close(outCh)
	} (chs, outCh)

	return outCh
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	out := mergeChannels(ch1, ch2)

	go func() {
		for m := range out {
			println(m)
		}
	}()

	ch1 <- 1
	ch2 <- 2
	ch1 <- 3

	time.Sleep(time.Millisecond)

	close(ch1)
	close(ch2)

	m, ok := <- ch1
	fmt.Print(m, ok)
}
