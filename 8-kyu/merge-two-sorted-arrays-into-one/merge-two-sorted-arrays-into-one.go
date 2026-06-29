package kata
​
import "sort"
​
func MergeArrays(arr1, arr2 []int) []int {
  var merged []int
  
  merged = append(merged, arr1...)
  merged = append(merged, arr2...)
  
  sort.Ints(merged)
  
 result := []int{}
  for i := 0; i < len(merged); i++ {
    if i == 0 || merged[i] != merged[i-1] {
     result = append(result, merged[i])
    }
  }
  return result
}
​