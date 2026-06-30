package kata
​
import (
  "strconv"
  "strings"
)
​
func HighAndLow(in string) string {
  numStr := strings.Fields(in)
  firstNum, _ := strconv.Atoi(numStr[0])
  max := firstNum
  min := firstNum
  for _, c := range numStr {
    num, _ := strconv.Atoi(c)
    if num > max {
      max = num
    }
    if num < min {
      min = num
    }
  }
  return strconv.Itoa(max) + " " + strconv.Itoa(min)
}