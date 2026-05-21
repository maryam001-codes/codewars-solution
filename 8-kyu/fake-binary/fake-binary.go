package kata
​
import "strconv"
​
func FakeBin(x string) string {
  var result string
  
  for _, v := range x{
    num, _ := strconv.Atoi(string(v))
    if num < 5 {
      result += "0"
    } else {
      result += "1"
    }
  }
  return result
}
​