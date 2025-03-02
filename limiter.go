package main

import (
	"fmt"
	"time"
)

func main() {
	lim := time.Tick(100 * time.Millisecond)

	for l := range lim {
		fmt.Println(l)
	}
}
