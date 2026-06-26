package kata
​
func CountRedBeads(n int) int {
  var res int
  
  if n < 2 {
    res = 0
  }
  
  if n == 2{
    res = 2
  }
  
  if n > 2 {
    res = (n * 2) - 2
  }
  return res
}