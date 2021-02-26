package main

import "fmt"

const BASE = 211

func Pow(x int) uint32 {
  ans := uint32(1)
  for i := 1; i <= x; i++ {
    ans *= BASE
  }
  return ans
}

func Hash(text string) uint32 {
  h := uint32(0)
  for _, bp := range text {
    h = h * BASE + uint32(bp)
  }
  return h
}

func PatternCount(text, pattern string) int {
  k := len(pattern)
  phash, thash := Hash(pattern), Hash(text[0:k])
  count, l , r := 0, 0, k
  power := Pow(k-1)
  for {
    if phash == thash {
      count++
    }
    if r == len(text) {
      break
    }
    thash = thash - (uint32(text[l]) * power)
    thash = (thash * BASE) + uint32(text[r])
    l++
    r++

  }
  return count
}



func main() {
  //test()
  var text, pattern string
  fmt.Scanln(&text)
  fmt.Scanln(&pattern)
  fmt.Println(PatternCount(text, pattern))
}
