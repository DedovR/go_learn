package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
  urls := []string{
    "a",
    "b",
    "c",
    "d",
    "e",
  }

  wg := &sync.WaitGroup{}
  doneCh := make(chan struct{})
  once := &sync.Once{}

  for _, url := range urls {

    wg.Add(1)
    go func(url string) {
      defer wg.Done()

      select {
      case d, ok := <-doneCh:
        fmt.Println("EVACUATION", d, ok)
        return
      default:
      }

      fmt.Println("fetching...")

      err := fetchUrl(url)
      if err != nil {
        fmt.Printf("Error fetching %s\n", url)

        once.Do(func(){
          close(doneCh)
        })
        return
      }

      fmt.Printf("Fetching %s\n", url)
    }(url)
  }

  wg.Wait()
  fmt.Println("success")
}

func fetchUrl(url string) error {
  r := rand.Int() % 5
  if r == 4 {
    return errors.New("ERROR")
  }
  time.Sleep(time.Second * time.Duration(r))
  return nil
}
