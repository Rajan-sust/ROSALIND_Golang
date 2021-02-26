package main

import "fmt"

func Pow(a, x int) int {
  ans := 1
  for i := 1; i <= x; i++ {
    ans *= a
  }
  return ans
}

func SeqToNum(kmer string) int {
  mp := map[byte]int{
    'A' : 0,
    'C' : 1,
    'G' : 2,
    'T' : 3,
  }
  num := 0
  for _, bp := range kmer {
    num = num * 4 + mp[byte(bp)]
  }
  return num
}

func FrequencyFind(text string, k int) []int {
  mp := map[string]int{}
  for i := 0; i < len(text) - k + 1; i++ {
    mp[text[i:i+k]]++
  }
  frequency := make([]int, Pow(4, k))
  for k, v := range mp {
    frequency[SeqToNum(k)] = v
  }
  return frequency
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
  // fmt.Print(SeqToNum("AAAA"))
  fmt.Scan(&text, &k)
  FormattedPrint(FrequencyFind(text, k))
}
