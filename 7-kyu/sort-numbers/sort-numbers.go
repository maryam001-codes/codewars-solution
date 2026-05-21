package kata
​
import "sort"
​
func SortNumbers(numbers []int) []int {
  if numbers == nil {
    return []int{}
  }
  sort.Ints(numbers)
  return numbers
}