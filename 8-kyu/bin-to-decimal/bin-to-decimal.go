package kata
​
import "strconv"
​
func BinToDec(bin string) int {
  dec, _ := strconv.ParseInt(bin, 2, 64)
  return int(dec)
}
​