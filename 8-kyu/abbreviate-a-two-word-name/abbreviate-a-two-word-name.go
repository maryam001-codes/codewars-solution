package kata
​
import (
  "fmt"
  "strings"
)
​
func AbbrevName(name string) string{
  
  names := strings.Fields(name)
  
  firstName := names[0]
  lastName := names[1]
  
  return fmt.Sprintf("%v.%v", strings.ToUpper(firstName[:1]), strings.ToUpper(lastName[:1]))
}