package main

import (
	"fmt"
	"time"
)

var m = map[string]int{"a": 1}

func main() {
  go read()
  time.Sleep(1 * time.Second)

  go write()
  // fatal error: concurrent map read and map write
  time.Sleep(1 * time.Second)
}

func read() {
  for {
    fmt.Println(m["a"])
  }
}

func write() {
  for {
    m["b"] = 42
  }
}
