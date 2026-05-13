package kata
​
import (
  "strconv"
  "strings"
)
​
​
func Points(games []string) int {
  var res int
  
  for _, s := range games {
     res += GetPoints(s)
  }
  return res
}
​
func GetPoints(score string) int{
  team := strings.Split(score, ":")
  x := team[0]
  y := team[1]
  
  i, _ := strconv.Atoi(x)
  j, _ := strconv.Atoi(y)
  
  if i > j {
    return 3
  } else if i < j {
    return 0
  }
  return 1
}