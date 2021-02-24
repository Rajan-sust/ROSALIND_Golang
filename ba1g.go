package main

import (
  "fmt"
  "RosalindGolang/utility"
)

func main() {
  var text1, text2 string
  fmt.Scan(&text1, &text2)
  fmt.Println(utility.HammingDist(text1, text2))
}
