package kata
​
import "unicode"
​
func ReverseLetters(s string) string {
  var output string
  var letters string
    for _, char := range s {
      if unicode.IsLetter(char){
        letters += string(char)
      }
    }
  for _, char := range letters{
    output = string(char) + output
  }
   return output
}
​