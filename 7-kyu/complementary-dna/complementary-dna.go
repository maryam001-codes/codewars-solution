package kata
​
func DNAStrand(dna string) string {
  var result string
  sub := map[rune]rune{
    'A' : 'T',
    'T' : 'A',
    'C' : 'G',
    'G' : 'C',
  }
  for _, ch := range dna {
    result += string(sub[ch])
  }
  return result
}
​
​