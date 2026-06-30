package kata
​
func Number(busStops [][2]int) int {
  currentNum := 0
  for _, stop := range busStops {
    currentNum += stop[0]
    currentNum -= stop[1]
  }
  return currentNum
}