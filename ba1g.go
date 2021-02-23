package main

import "fmt"

func HammingDist(text1, text2 string) int {
  d := 0
  for i := 0; i < len(text1); i++ {
    if text1[i] != text2[i] {
      d++
    }
  }
  return d
}

func main() {
  var text1, text2 string
  fmt.Scan(&text1, &text2)
  fmt.Println(HammingDist(text1, text2))
}
