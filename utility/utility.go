package utility

func HammingDist(text1, text2 string) int {
  d := 0
  for i := 0; i < len(text1); i++ {
    if text1[i] != text2[i] {
      d++
    }
  }
  return d
}
