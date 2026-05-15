package kata
​
func Maps(x []int) []int {
  res := make([]int, len(x))
  
  for i, v := range x{
    res[i] = v * 2
  }
  return res
}