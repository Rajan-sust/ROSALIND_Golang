package main

import (
  "fmt"
  "RosalindGolang/utility"
)

func FrequencyFind(text string, k int) []int {
  mp := map[string]int{}
  for i := 0; i < len(text) - k + 1; i++ {
    mp[text[i:i+k]]++
  }
  arr := utility.GenerateKmer(k)
  fre := []int{}
  for _, kmer := range(arr) {
    if val, ok := mp[kmer]; ok {
      fre = append(fre, val)
    } else {
      fre = append(fre, val)
    }
  }
  return fre
}

func FormattedPrint(frequency []int) {
  ok := true
  for _, val := range frequency {
    if ok {
      fmt.Print(val)
      ok = false
    } else {
      fmt.Print(" ", val)
    }
  }
  fmt.Println()
}

func main() {
  var text string
  var k int
  fmt.Scan(&text, &k)
  FormattedPrint(FrequencyFind(text, k))
}
