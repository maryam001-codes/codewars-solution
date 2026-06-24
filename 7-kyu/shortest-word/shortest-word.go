package kata
​
import "strings"
​
func FindShort(s string) int {
  words := strings.Fields(s)
  
  minLength := len(words[0])
  
  for _, word := range words {
    if len(word) < minLength {
      minLength = len(word)
    }
  }
  return minLength
}