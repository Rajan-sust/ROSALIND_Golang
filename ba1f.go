package main

import (
  "fmt"
)


func FindMinSkew(text string) (pos []int) {
  MinSkew, skew := 0, 0

  for p, b := range text {
    switch b {
      case 'G':
        skew++
      case 'C':
        skew--
    }
    if skew == MinSkew {
      pos = append(pos, p + 1)
    } else if skew < MinSkew {
      MinSkew = skew
      pos = []int{p + 1}
    }
  }
  return
}

func main() {
  var text string
  fmt.Scan(&text)
  for _, pos := range FindMinSkew(text) {
    fmt.Print(pos, " ")
  }
}
