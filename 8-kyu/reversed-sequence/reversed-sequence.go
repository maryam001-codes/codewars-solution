package kata
​
func ReverseSeq(n int) []int {
  var result []int
  num := 1
  for i := n-1; i >= 0 && i < n; i-- {
    if num + i <= n{
      result = append(result, num+i)
    }
  }
  return result
}