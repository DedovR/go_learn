package main

import (
	"fmt"
	"context"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	numChan := make(chan int)

	go work(ctx, numChan)

	go func() {
		for i := 0; i < 5; i++ {
			numChan <- i
		}
	}()

	cancelFunc()

	time.Sleep(1 * time.Second)
}

func work(ctx context.Context, numChan chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop")
			return
		case num := <-numChan:
			fmt.Println(num)
		}
	}
}
