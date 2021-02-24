package main

import (
  "fmt"
  "RosalindGolang/utility"
)

func ApproximateMatch(pattern, text string, d int) (pos []int) {
  for i:= 0; i < len(text) - len(pattern) + 1; i++ {
    if utility.HammingDist(text[i:i+len(pattern)], pattern) <= d {
      pos = append(pos, i)
    }
  }
  return
}

func main() {
  var pattern, text string
  var d int
  fmt.Scan(&pattern, &text, &d)
  for _, v := range ApproximateMatch(pattern, text, d) {
    fmt.Print(v, " ")
  }
}
