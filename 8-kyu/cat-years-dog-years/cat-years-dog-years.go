package kata
​
func CalculateYears(years int) (result [3]int) {
  var catYears, dogYears int
  switch years{
    case 1:
    catYears = 15
    dogYears = 15
    case 2:
    catYears= 15 + 9
    dogYears= 15 + 9
    default:
    catYears = 24 + (years-2)*4
    dogYears = 24 + (years-2)*5
  }
  return [3]int{years, catYears, dogYears} 
}