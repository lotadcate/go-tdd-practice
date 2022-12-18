package arrayandslices

func SumAllTails(arrToSum ...[]int) []int {
  var sums []int
  for _, numbers := range arrToSum {
    if len(numbers) == 0 {
      sums = append(sums, 0)
    } else {
      tail := numbers[1:]
      sums = append(sums, Sum(tail))
    }
  }
  return sums
}
