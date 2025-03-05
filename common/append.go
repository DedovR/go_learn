package main

import "fmt"

func a() {
  x := []int{}
  x = append(x, 0)  // [0], len=1, cap=1
  x = append(x, 1)  // [0, 1}, len=2, cap=2
  x = append(x, 2)  // [0, 1, 2], len=3, cap=4
  y := append(x, 3) // [0, 1, 2], len=3, cap=4
  z := append(x, 4) // [0, 1, 2], len=3, cap=4
  print(x)          // [0, 1, 2], len=3, cap=4
  fmt.Println(y, z)
  // [0,1,2,4] [0,1,2,4] так как `y` ссылается на туже память что и z
}

func foo(src []int) {
  src = append(src, 5)
}

func print(a []int) {
  fmt.Printf("%v len=%d cap=%d ptr=%p\n", a, len(a), cap(a), &a)
}

func main() {
	// a()

  arr := []int{1, 2, 3}
  src := arr[:1]
  // src := make([]int, 1, len(arr))
  // copy(src, arr[:1])   // [1] len=2 cap=3 - второй ноль, потому что в функцию уходит по значению и лен не апдейтится

  print(arr)
  foo(src)
  arr = append(arr, 4)
  print(src) // [1] len=1 cap=3
  print(arr) // [1 5 3] len=3 cap=3
}
