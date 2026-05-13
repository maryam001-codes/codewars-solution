package kata
//12345
func Grow(arr []int) int{
  res := 1
  for _, c := range arr{
    res = res * c
  }
  return res
  
}