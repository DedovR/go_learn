package main

import (
	"fmt"
	"time"
)

func main() {
  ch := make(chan int, 1)
  go func() {
    ch <- 1
  }()

  time.Sleep(1 * time.Second)
  close(ch)

  for i := range ch {
    fmt.Println(i)
  }
}
