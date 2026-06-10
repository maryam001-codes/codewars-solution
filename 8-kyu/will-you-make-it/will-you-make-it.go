package kata
​
func ZeroFuel(distanceToPump int, mpg int, fuelLeft int) bool {
  return distanceToPump <= mpg * fuelLeft
}