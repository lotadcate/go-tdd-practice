package iteration

// const repeatCount = 5
func Repeat(c string, count int) string {
  var repeated string
  for i:=0; i<count; i++ {
    repeated += c
  }
  return repeated
}
