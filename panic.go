package main

import (
	"fmt"
	"time"
)

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    go g(0)
		time.Sleep(5 * time.Second)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
		defer func() {
			if r := recover(); r != nil {
					fmt.Println("Recovered in g", r)
			}
		}()
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i+1)
}
