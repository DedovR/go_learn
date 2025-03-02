package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
  n, err := r13.r.Read(b)
    for i := 0; i <= n; i++ {
        b[i] = b[i] + 13
    }
    return n, err
}

func main() {
  s := strings.NewReader("Make golang great again")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
