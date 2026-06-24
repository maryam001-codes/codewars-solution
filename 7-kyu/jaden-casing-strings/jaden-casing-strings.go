package kata
​
import "strings"
​
func ToJadenCase(str string) string {
  words := strings.Fields(str)
  
  for i, word := range words {
    if len(word) > 0 {
     words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:]) 
    }
  }
  return strings.Join(words, " ")
}