package main

import (
  "fmt"
  "strings"
)

func Max(x, y int) int {
  if x > y {
    return x
  }
  return y
}

func FrequentWord(text string, k int) (kmers []string) {
  mx, mp := 0, make(map[string]int)
  for i := 0; i < len(text) - k + 1; i++ {
    mp[text[i:i+k]]++
    mx = Max(mp[text[i:i+k]], mx)
  }

  for k, v := range mp {
    if v == mx {
      kmers = append(kmers, k)
    }
  }
  return
}

func main()  {
  var text string
  var k int
  fmt.Scan(&text, &k)
  fmt.Println(strings.Join(FrequentWord(text, k), " "))

}
