package kata
​
func FindOutlier(integers []int) int {
  var even []int
  var odd []int
  
  for _, c := range integers {
    if c&1 == 0 {
      even = append(even, c)
    } else {
      odd = append(odd, c)
      }
    }
  
    if len(even) == 1{
      return even[0]
      }
  return odd[0]
  }