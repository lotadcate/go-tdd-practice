package arrayandslices

// Takes a number of different slices and returns a new slice containing the SumAll of each slice passed.
func SumAll(arrtoSum ...[]int) []int {
  lengthOfArr := len(arrtoSum)
  sums := make([]int, lengthOfArr)

  for i, arrs := range arrtoSum {
    sums[i] = Sum(arrs)
  }
  return sums
}
