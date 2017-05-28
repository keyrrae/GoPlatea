package main

import (
  "fmt"
  "flag"
  "strconv"
  "math/rand"
)

type Matrix [][]float64

func Multiply(m1, m2 Matrix) (m3 Matrix, ok bool) {
  rows, cols, extra := len(m1), len(m2[0]), len(m2)
  if len(m1[0]) != extra {
      return nil, false
  }
  m3 = make(Matrix, rows)
  for i := 0; i < rows; i++ {
      m3[i] = make([]float64, cols)
      for j := 0; j < cols; j++ {
          for k := 0; k < extra; k++ {
              m3[i][j] += m1[i][k] * m2[k][j]
          }
      }
  }
  return m3, true
}

func (m Matrix) String() string {
  rows := len(m)
  cols := len(m[0])
  out := "["
  for r := 0; r < rows; r++ {
      if r > 0 {
          out += ",\n "
      }
      out += "[ "
      for c := 0; c < cols; c++ {
          if c > 0 {
              out += ", "
          }
          out += fmt.Sprintf("%7.3f", m[r][c])
      }
      out += " ]"
  }
  out += "]"
  return out
}

func main() {
  flag.Parse()
  n := 100
  if flag.NArg() > 0 {
     size, _ := strconv.Atoi(flag.Arg(0))
     n = size
  }
  var A Matrix
  var B Matrix
  for i:=0; i < n; i++ {
    var a []float64
    var b []float64
    for j:=0; j < n; j++ {
      a = append(a, rand.Float64())
      b = append(b, rand.Float64())
    }
    A = append(A, a)
    B = append(B, b)
  }
  _, ok := Multiply(A, B)
  if !ok {
      panic("Invalid dimensions")
  }
}
