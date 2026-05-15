package kata
​
import "unicode"
​
func ToAlternatingCase(str string) string {
  var res string
​
  for _, char := range str {
    if char >= 'A' && char <= 'Z' {
      char = unicode.ToLower(char)
    } else if char >= 'a' && char <= 'z' {
      char = unicode.ToUpper(char)
    }
    res += string(char)
  }
  return res
}