package arrayandslices

func Sum(arr [5]int) int {
  ans := 0
  for i:=0; i<5; i++ {
    ans += arr[i]
  }
  return ans
}
