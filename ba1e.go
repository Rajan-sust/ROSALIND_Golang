/*
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
*/

package main

import (
  "fmt"
  "strings"
)

type Pair struct {
  X int
  Y int
}


func ClumpFinder(text string, k, L, t int) []string {
  mp := map[string][]Pair{}
  for i := 0; i < len(text) - k + 1; i++ {
    kmer := text[i:i+k]
    mp[kmer] = append(mp[kmer], Pair{i, i + k - 1})
  }
  kmers := []string{}
  for key, val := range mp {
    n := len(val)
    if n >= t {
      for j := 0; j < n - t + 1; j++ {
        start := val[j].X
        finish := val[j + t -1].Y
        if finish - start + 1 <= L {
          kmers = append(kmers, key)
          break
        }
      }
    }
  }
  return kmers
}


func main() {
  var text string
  var k, L, t int
  fmt.Scan(&text, &k, &L, &t)
  fmt.Println(strings.Join(ClumpFinder(text, k, L, t), " "))
}
