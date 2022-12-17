package arrayandslices

func Sum(arr [5]int) int {
  ans := 0
  for _, number := range arr {
    ans += number
  }
  return ans
}
