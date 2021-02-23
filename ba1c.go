package main

import "fmt"

func ReverseComplement(text string) string {
  mp := map[byte]byte{'A': 'T', 'T': 'A', 'C': 'G', 'G': 'C'}
  var rc []byte
  for i := len(text) - 1; i >= 0; i-- {
    rc = append(rc, mp[text[i]])
  }
  return string(rc)
}

func main() {
  var text string
  fmt.Scan(&text)
  fmt.Println(ReverseComplement(text))
}
