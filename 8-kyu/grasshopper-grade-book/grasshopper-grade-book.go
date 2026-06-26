package kata
​
func GetGrade(a,b,c int) rune {
  average := (a+b+c) / 3
  switch {
    case average >= 90 && average <= 100:
      return 'A'
    case average >= 80 && average < 90:
      return 'B'
    case average >= 70 && average < 80:
      return 'C'
    case average >= 60 && average < 70:
      return 'D'
    default :
    return 'F'
    }
}