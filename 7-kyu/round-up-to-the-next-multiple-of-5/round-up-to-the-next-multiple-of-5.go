package kata
​
func RoundToNext5(n int) int {
  result := 0
  remainder := n % 5 
  if remainder > 0 {
    result = n + (5 - remainder)
  } else if  remainder < 0 {
    result = n - remainder
  } else if remainder == 0 {
    return n
  }
  return result
}
​