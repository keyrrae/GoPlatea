package main

import (
  "os"
  "fmt"
)

func getCount() int {
  count := 0
  for i := 0; i < 1000; i++ {
    count++
  }
  return count
}

func main() {
  fmt.Println(os.Args)
}
