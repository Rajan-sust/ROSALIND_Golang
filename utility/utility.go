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

func GenerateKmer(k int) (arr []string) {
  kmers := []string{"A", "C", "G", "T"}
  previous := 0
  for i := 1; i <= k - 1; i++ {
    n := len(kmers)
    for j := previous; j < n; j++ {
      for _, bp := range []string{"A", "C", "G", "T"} {
        kmers = append(kmers, kmers[j] + bp)
        if i == k - 1 {
          arr = append(arr, kmers[j] + bp)
        }
      }
    }
    previous = n
  }
  return
}
