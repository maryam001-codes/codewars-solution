package kata
​
import "sort"
​
func TwoSort(arr []string) string {
  var res string
  
  if len(arr) <= 0 || arr == nil{
    return ""
  }
  
  sort.Strings(arr)
  
  firstValue := arr[0]
  for i, char := range firstValue{
    if i != len(firstValue) - 1 {
      res += string(char) + "***"
    } else {
    res += string(char)
      }
    }
  return res
}