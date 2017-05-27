package main

import (
  "os"
//  "fmt"
  "strconv"
)

func getCount() int {
  count := 0
  for i := 0; i < 1000; i++ {
    count++
  }
  return count
}

func main() {
  //fmt.Println(os.Args)
  n, _ := strconv.Atoi(os.Args[1])
  for i:= 0; i < n; i++ {
    getCount()
  }
}
