package main

import "fmt"

func a() {
  x := []int{}
  x = append(x, 0)  // [0], len=1, cap=1
  x = append(x, 1)  // [0, 1}, len=2, cap=2
  x = append(x, 2)  // [0, 1, 2], len=3, cap=4
  y := append(x, 3) // [0, 1, 2], len=3, cap=4
  z := append(x, 4) // [0, 1, 2], len=3, cap=4
  fmt.Println(x, len(x), cap(x)) // [0, 1, 2], len=3, cap=4
  fmt.Println(y, z)
  // [0,1,2,4] [0,1,2,4] так как `y` ссылается на туже память что и z
  x = append(x, 3)
  p := &x[3]
  fmt.Println(x, len(x), cap(x))
  fmt.Println(p)
  x = append(x, 4)
  fmt.Println(x, len(x), cap(x))
  fmt.Println(p, &x[3])
}

func main() {
	a()
}
