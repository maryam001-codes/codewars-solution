package kata
​
import "strings"
​
func ToNato(words string) string {
  
  // the NATO map[string]string is preloaded for you
// NATO["A"] == "Alfa", etc
  
  var result []string
  
  for _, ch := range words{
    charStr := strings.ToUpper(string(ch))
    
    if charStr == " " {
      continue
    }
    if natoWord, exists := NATO[charStr]; exists {
      result = append(result, natoWord)
    } else {
      result = append(result, charStr)
    }
  }
  return strings.Join(result, " ")
}
​
​