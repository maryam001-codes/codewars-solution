package kata
​
func BoolToWord(word bool) string {
  words := map[bool]string{true:"Yes", false:"No",}
  return words[word]
}