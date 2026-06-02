package kata
​
​
func SequenceSum(start, end, step int) int {
  var result int
  for i := start; i <= end; i += step{
    result += i
  }
  return result
}
​