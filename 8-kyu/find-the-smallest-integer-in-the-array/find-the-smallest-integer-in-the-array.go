package kata
​
func SmallestIntegerFinder(numbers []int) int {
  min := numbers[0]
  
  for _, num := range numbers{
    if num < min {
      min=num
    }
  }
  return min
}