package main

import "fmt"

type MyError struct {
  data string
}

func (e MyError) Error() string {
  return e.data
}

func main() {
  err := foo(4)
  if err != nil {
    fmt.Println("oops")
  } else {
    fmt.Println("ok")
  }
}

func foo(a int) error {
  var err *MyError
  if a > 5 {
    err = &MyError{data: "i > 5"}
  }

  return err
}
