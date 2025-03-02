package main

import "fmt"

func main() {
	defer func() {
	 if panicObj := recover(); panicObj != nil {
		fmt.Printf("case D\n")
	 }
	}()

	ch1 := make(chan int, 1)
	ch2 := make(chan string)

	close(ch1)
	close(ch2)

	select {
	case <-ch2:
	 fmt.Printf("case B\n") // или этот, так как будем получать zero-value
	case ch1<-0: // или этот, но будет паника при попытке писать в закрытый канал, панику перехватит дефер и выведет Д
	 fmt.Printf("case A\n")
	default:
	 fmt.Printf("case C\n")
	}
 }
