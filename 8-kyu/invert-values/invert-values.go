package kata
​
func Invert(arr []int) []int {
  var result []int
  for _, num := range arr {
    num = -num
    result = append(result, num)
  }
  return result
}
​