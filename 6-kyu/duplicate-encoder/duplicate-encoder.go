package kata
​
import "strings"
​
func DuplicateEncode(word string) string {
  word = strings.ToLower(word)
  
  freq := make(map[rune]int)
  
  for _, char := range word {
    freq[char]++
  }
  
  var result strings.Builder
  
  for _, char := range word {
    if freq[char] > 1 {
      result.WriteString(")")
    } else {
    result.WriteString("(")
      }
    }
  return result.String()
}