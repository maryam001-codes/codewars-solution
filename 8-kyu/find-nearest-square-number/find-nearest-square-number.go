package kata
​
func NearestSq(n int) int {
  i := 1
  
  for i*i < n{
    i++
  }
  
  upper := i*i
  lower := (i-1) * (i-1)
  
  if upper - n > n - lower {
    return lower
  }
  return upper
}