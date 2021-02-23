package main

import "fmt"

func FindPosition(pattern, text string) (pos []int) {
  k := len(pattern)
  for i := 0; i < len(text) - k + 1; i++ {
    if text[i:i+k] == pattern {
      pos = append(pos, i)
    }
  }
  return pos
}

func main() {
  var pattern, text string
  fmt.Scan(&pattern, &text)
  flag := false
  for _, p := range FindPosition(pattern, text) {
    if !flag {
      fmt.Print(p)
      flag = true
    } else {
      fmt.Print(" ", p)
    }
  }
}
