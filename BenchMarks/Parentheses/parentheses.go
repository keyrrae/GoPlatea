package main

import (
  "math/rand"
  "fmt"
)

func removeInvalidParentheses(str string) []string {
  var res []string
  if str == "" {
    return res
  }

  visited := make(map[string]bool)
  var queue []string

  queue = append(queue, str)
  visited[str] = true

  found := false;
  for len(queue) != 0 {
    s := queue[0]
    //fmt.Println(s)
    queue = queue[1:]
    if isValid(s) {
      //fmt.Println("valid")
      res = append(res, s)
      found = true
    }
    if found {
      continue
    }
    ss := []byte(s)
    for i := 0; i < len(ss); i++ {
      if ss[i] != '(' && ss[i] != ')' {
        continue
      }

      t := append(ss[:i], ss[i+1:]...)
      tt := string(t)
      if _, ok := visited[tt]; !ok {
        queue = append(queue, tt)
        visited[tt] = true
      }
    }
  }
  return res
}

func isValid(str string) bool {
  count := 0
  s := []byte(str)
  for i := 0; i < len(s); i++ {
    c := s[i]
    if c == '(' {
      count++
    }
    if c == ')' {
      count--
      if count == -1 {
        return false
      }
    }
  }
  return count == 0
}

func main() {
  s := ""
  n := 40
  par := []byte("(()()(()(()((()))(()((())(()(()((()())((())))()")
  rand.Seed(1243)
  for i := 0; i < n; i++ {
    s = s + string(par[ i % len(par)])
  }
  fmt.Println(s)

  fmt.Println(removeInvalidParentheses(s))
}
