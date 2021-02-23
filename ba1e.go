package main

import (
  "fmt"
  "strings"
)

func AnyKmer(text string, k, t int) (kmers []string) {
  mp := map[string]int{}
  for i := 0; i < len(text) - k + 1; i++ {
    mp[text[i:i+k]]++
    if mp[text[i:i+k]] >= t {
      kmers = append(kmers, text[i:i+k])
    }
  }
  return
}

func FindClump(text string, k, L, t int) (kmers []string) {
  mp := map[string]bool{}
  for i := 0; i < len(text) - L + 1; i++ {
      for _, v := range AnyKmer(text[i:i+L], k, t) {
        mp[v] = true
      }
  }
  for key, _ := range mp {
    kmers = append(kmers, key)
  }
  return
}

func main() {
  var text string
  var k, L, t int
  fmt.Scan(&text, &k, &L, &t)
  fmt.Println(strings.Join(FindClump(text, k, L, t), " "))
}
