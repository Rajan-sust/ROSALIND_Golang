package main

import "fmt"

func PatternCount(text, pattern string) (count int) {
  for i := 0; i < len(text) - len(pattern) + 1; i++ {
    if text[i:i+len(pattern)] == pattern {
      count++
    }
  }
  return
}

func main() {
  var text, pattern string
  fmt.Scanln(&text)
  fmt.Scanln(&pattern)
  fmt.Println(PatternCount(text, pattern))
}
