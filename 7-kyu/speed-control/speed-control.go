package kata
​
func Gps(s int, x []float64) int {
  if len(x) <= 1 {
    return 0
  }
  maxSpeed := 0.0
  for i := 1; i < len(x); i++{
    deltaDistance := x[i] - x[i-1]
    currentSpeed := (3600.0 * deltaDistance) / float64(s)
    if currentSpeed > maxSpeed {
      maxSpeed = currentSpeed
    }
  }
  return int(maxSpeed)
}