package kata
​
func HighestRank(nums []int) int {
  count := make(map[int]int)
  
  for _, num := range nums {
    count[num]++
  }
  
  maxFreq := 0
  result := nums[0]
  
  for num, freq := range count{
    if freq > maxFreq || (freq == maxFreq && num > result) {
      maxFreq = freq
      result = num
    }
  }
  return result
}
​