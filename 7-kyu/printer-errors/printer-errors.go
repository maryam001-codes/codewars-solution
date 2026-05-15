package kata
​
import "fmt"
​
func PrinterError(s string) string {
  length := len(s)
  count := 0
  
  for _, char := range s {
    if char < 'a' || char > 'm' {
      count ++
    }
  } 
  return fmt.Sprintf("%v/%v", count, length)
}