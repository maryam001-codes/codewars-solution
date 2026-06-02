package kata
​
func Summation(n int) int {
  var result int
  for i := 1; i <=n; i++ {
    result += i
  }
  return result
}