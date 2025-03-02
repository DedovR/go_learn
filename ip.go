package main

import (
  "fmt"
  "strings"
  "strconv"
)

type IPAddr [4]byte

func (t IPAddr) String() string {
  var sb strings.Builder
  for i, a := range t {
    sb.WriteString(strconv.Itoa(int(a)))
    if i < 3 {
      sb.WriteString(".")
    }
  }
  return sb.String()
}

func main() {
  hosts := map[string]IPAddr{
    "loopback":  {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }
  for name, ip := range hosts {
    fmt.Printf("%v: %v\n", name, ip)
  }
}
