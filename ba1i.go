package main

import (
  "fmt"
  "RosalindGolang/utility"
)

func MostFrequent(text string, k, d int) (kmers []string) {
  mx, mp := 0, map[string]int{}

  for _, kmer := range utility.GenerateKmer(k) {
    for i := 0; i < len(text) - k + 1; i++ {
      if utility.HammingDist(text[i:i+k], kmer) <= d {
        mp[kmer]++
      }
    }
    if mp[kmer] == mx {
      kmers = append(kmers, kmer)
    } else if mp[kmer] > mx {
      kmers = []string{kmer}
      mx = mp[kmer]
    }
  }
  return
}

func main() {
  var text string
  var k, d int
  fmt.Scan(&text, &k, &d)
  for _, kmer := range MostFrequent(text, k, d) {
    fmt.Print(kmer, " ")
  }
}
