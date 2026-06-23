package kata
​
func solve(str string) int {
  currentSum := 0
  maxSum := 0
  for _, char := range str {
    switch char {
      case 'a', 'e', 'i', 'o', 'u' :
      if currentSum > maxSum {
        maxSum = currentSum
      }
      currentSum = 0
      default :
      charValue := int(char - 'a' + 1)
      currentSum += charValue
      }
    }
   if currentSum > maxSum {
      maxSum = currentSum
  }
  return maxSum
}