package main

import "fmt"

func fibonacci() func() int {
  n := 0
  prev1 := 0
  prev2 := 1
  return func () int {
		if n == 0 {
			n++
			return 0
		}
    prev1, prev2 = prev2, prev1 + prev2
    n++
    return prev1
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}

